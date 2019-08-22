package rpc

import (
	"db-client-api/entity"
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
	return nil, nil
}

/**
查询数据，单条记录
*/
func (this Db) QueryRow(item *entity.SqlMessage) (map[string]string, error) {
	return nil, nil
}

/**
查询记录，一个字段
*/
func (this Db) QueryScalar(item *entity.SqlMessage) (string, error) {
	return "", nil
}

/**
更新数据
*/
func (this Db) Update(item *entity.SqlMessage) (int64, error) {
	return 0, nil
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
	return 0, nil
}

/**
事务处理，多条SQL必须是操作同一数据库
SQL: UPDATE、INSERT、DELETE
*/
func (this Db) Transaction(items []*entity.SqlMessage) (bool, error) {
	return false, nil
}

/**
事务处理，多条SQL必须是操作同一数据库
SQL: UPDATE、INSERT、DELETE
*/
func (this Db) TransactionV2(items *entity.BatchSqlMessage) (bool, error) {
	return false, nil
}

/**
批量查询
*/
func (this Db) BatchQuery(items []*entity.SqlMessage) ([][]map[string]string, error) {
	return nil, nil
}
