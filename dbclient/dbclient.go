package dbclient

import (
	"github.com/sillyhatxu/sqlite-client"
)

var Client *client.SqliteClient

func InitialDBClient(dataSourceName string, ddlPath string) {
	Client = client.NewSqliteClient(dataSourceName, client.DDLPath(ddlPath))
	err := Client.Initial()
	if err != nil {
		panic(err)
	}
}
