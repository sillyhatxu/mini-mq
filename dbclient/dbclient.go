package dbclient

import "github.com/sillyhatxu/mysql-client"

var Client *dbclient.MysqlClient

func InitialDBClient(dataSourceName string, ddlPath string) {
	Client = dbclient.NewMysqlClientConf(dataSourceName, dbclient.DDLPath(ddlPath))
	err := Client.Initial()
	if err != nil {
		panic(err)
	}
}
