
数据库操作客户端
---

使用说明

```
// 服务器端使用，请查看 https://github.com/gentwolf-shen/db-service-web-api

// 安装
go get -u github.com/gentwolf-shen/db-client-api
```

```go

package main

import (
	"fmt"

	"github.com/gentwolf-shen/db-client-api/entity"
	"github.com/gentwolf-shen/db-client-api/web"
)

func main() {
	// 初始化
	dbServer := "http://127.0.0.1:9001"
	appKey := "bcebfefa29ca565acf9dfe17f0c6e863"
	appSecret := "4b654b10b12d0a72aee34dd11ed365c0"

	dbClient := web.New(dbServer, appKey, appSecret)

	testQuery(dbClient)
	testUpdate(dbClient)
	testBatchQuery(dbClient)
	testTransactionV1(dbClient)
	testTransactionV2(dbClient)
}

// 查询, Query(多条记录) / QueryRow(一条记录) / QueryScalar(首个字段值)
func testQuery(dbClient *web.Db) {
	item := &entity.SqlMessage{}
	item.Sql = "SELECT * from user WHERE id>=?"
	item.Params = []interface{}{2}

	// Query / QueryRow / QueryScalar
	rows, err := dbClient.Query(item)
	fmt.Println(err)
	fmt.Println(rows)
}

// 更新/删除, 返回影响的记录数
func testUpdate(dbClient *web.Db) {
	item := &entity.SqlMessage{}
	item.Sql = "INSERT INTO user set username=?, email=?"
	item.Params = []interface{}{"test", "test@test.com"}

	n, err := dbClient.Insert(item)
	fmt.Println(err)
	fmt.Println(n)
}

// 批量查询, 多维数组
func testBatchQuery(dbClient *web.Db) {
	item1 := &entity.SqlMessage{}
	item1.Sql = "SELECT * FROM user WHERE id=?"
	item1.Params = []interface{}{10}

	item2 := &entity.SqlMessage{}
	item2.Sql = "SELECT * FROM user WHERE id=?"
	item2.Params = []interface{}{11}

	items := make([]*entity.SqlMessage, 2)
	items[0] = item1
	items[1] = item2

	rows, err := dbClient.BatchQuery(items)
	fmt.Println(err)
	fmt.Println(rows)
}

// 事务处理, 不同的SQL
func testTransactionV1(dbClient *web.Db) {
	item1 := &entity.SqlMessage{}
	item1.Sql = "INSERT INTO user SET username=?,email=?"
	item1.Params = []interface{}{"username 1", "email1"}

	item2 := &entity.SqlMessage{}
	item2.Sql = "INSERT INTO user SET username=?,email=?"
	item2.Params = []interface{}{"username 2", "email2"}

	items := make([]*entity.SqlMessage, 2)
	items[0] = item1
	items[1] = item2

	bl, err := dbClient.TransactionV1(items)
	fmt.Println(err)
	fmt.Println(bl)
}

// 事务处理, 相同的SQL, 不同的参数
func testTransactionV2(dbClient *web.Db) {
	item := &entity.BatchSqlMessage{}
	item.Sql = "INSERT INTO user SET username=?,email=?"

	item.Params = [][]interface{}{
		{"username A", "email A"},
		{"username B", "email B"},
		{"username C", "email C"},
	}

	bl, err := dbClient.Transaction(item)
	fmt.Println(err)
	fmt.Println(bl)
}


```