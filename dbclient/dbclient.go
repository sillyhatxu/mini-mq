package dbclient

import "github.com/sillyhatxu/sqlite-client"

var Client *sqliteclient.SqliteClient

func InitialDBClient(dataSourceName string, ddlPath string) {
	Client = sqliteclient.NewSqliteClient(dataSourceName, ddlPath)
	err := Client.Initial()
	if err != nil {
		panic(err)
	}
}
