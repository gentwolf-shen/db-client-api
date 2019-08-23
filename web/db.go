package web

import (
	"encoding/json"
	"errors"

	"github.com/gentwolf-shen/db-client-api/auth"
	"github.com/gentwolf-shen/db-client-api/entity"

	"github.com/gentwolf-shen/gohelper/httpclient"
)

type Db struct {
	dbServer  string
	appKey    string
	appSecret string
}

func New(dbServer, appKey, appSecret string) *Db {
	return &Db{
		dbServer:  dbServer,
		appKey:    appKey,
		appSecret: appSecret,
	}
}

/**
查询数据，返回多条记录
*/
func (this Db) Query(item *entity.SqlMessage) ([]map[string]string, error) {
	b, err := this.send("/query", item)
	if err != nil {
		return nil, err
	}

	var rows []map[string]string
	err = json.Unmarshal(b, &rows)

	return rows, err
}

/**
查询数据，单条记录
*/
func (this Db) QueryRow(item *entity.SqlMessage) (map[string]string, error) {
	rows, err := this.Query(item)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, nil
}

/**
查询记录，一个字段
*/
func (this Db) QueryScalar(item *entity.SqlMessage) (string, error) {
	row, err := this.QueryRow(item)
	if err != nil {
		return "", err
	}

	value := ""

	if len(row) > 0 {
		for _, val := range row {
			value = val
			break
		}
	}

	return value, nil
}

/**
更新数据
*/
func (this Db) Update(item *entity.SqlMessage) (int64, error) {
	b, err := this.send("/update", item)
	if err != nil {
		return 0, err
	}

	rs := &entity.UpdateResult{}
	if err := json.Unmarshal(b, rs); err != nil {
		return 0, err
	}

	return rs.AffectedRows, nil
}

/**
删除数据
*/
func (this Db) Delete(item *entity.SqlMessage) (int64, error) {
	return this.Update(item)
}

/**
添加数据
*/
func (this Db) Insert(item *entity.SqlMessage) (int64, error) {
	b, err := this.send("/insert", item)
	if err != nil {
		return 0, err
	}

	rs := &entity.InsertResult{}
	if err := json.Unmarshal(b, rs); err != nil {
		return 0, err
	}

	return rs.LastInsertId, nil
}

/**
事务处理，多条SQL必须是操作同一数据库
SQL: UPDATE、INSERT、DELETE
*/
func (this Db) Transaction(items []*entity.SqlMessage) (bool, error) {
	if _, err := this.send("/v1/transaction", items); err != nil {
		return false, err
	}

	return true, nil
}

/**
事务处理，多条SQL必须是操作同一数据库
SQL: UPDATE、INSERT、DELETE
*/
func (this Db) TransactionV2(items *entity.BatchSqlMessage) (bool, error) {
	if _, err := this.send("/v2/transaction", items); err != nil {
		return false, err
	}

	return true, nil
}

/**
批量查询
*/
func (this Db) BatchQuery(items []*entity.SqlMessage) ([][]map[string]string, error) {
	b, err := this.send("/batch/query", items)
	if err != nil {
		return nil, err
	}

	var rows [][]map[string]string
	err = json.Unmarshal(b, &rows)

	return rows, err
}

/**
发送数据操作命令
*/
func (this Db) send(method string, item interface{}) ([]byte, error) {
	token, err := auth.GetToken(this.appKey, this.appSecret)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string, 3)
	headers["Authorization"] = this.appKey + ":" + token
	headers["Connection"] = "keep-alive"

	b, _ := json.Marshal(item)

	body, err := httpclient.PostToBody(this.dbServer+method, b, headers)
	if err != nil {
		return nil, errors.New(string(body))
	}

	return body, nil
}
