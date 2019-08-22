
数据库操作客户端
---

服务器:
https://github.com/gentwolf-shen/db-service-web-api

使用

```

// 1. 初始化
dbClient := web.New("http://127.0.0.1:9001", "bcebfefa29ca565acf9dfe17f0c6e863", "4b654b10b12d0a72aee34dd11ed365c0")

// 2. 调用
item := &entity.SqlMessage{}
item.Sql = "SELECT * from user WHERE id>=?"
item.Params = []interface{}{2}

rows, err := dbClient.Default.Query(item)
	
```