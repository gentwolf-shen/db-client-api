package entity

type (
	SqlItem struct {
		Sql string `json:"sql" binding:"required"`
	}

	SqlMessage struct {
		SqlItem

		Params []interface{} `json:"params"`
	}

	BatchSqlMessage struct {
		SqlItem

		Params [][]interface{} `json:"params"`
	}

	UpdateResult struct {
		AffectedRows int64 `json:"affectedRows"`
	}

	InsertResult struct {
		LastInsertId int64 `json:"lastInsertId"`
	}
)
