package main

import (
	"github.com/sillyhatxu/mini-mq/api"
	"github.com/sillyhatxu/mini-mq/dbclient"
)

func main() {
	dbclient.InitialDBClient("./basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	err := api.InitialAPI()
	if err != nil {
		panic(err)
	}
}
