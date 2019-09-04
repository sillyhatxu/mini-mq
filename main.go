package main

import (
	"github.com/sillyhatxu/mini-mq/api"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/grpcserver"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"net"
)

func main() {
	dbclient.InitialDBClient("./basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	cache.Initial()
	apiListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	grpcListener, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	go api.InitialAPI(apiListener)
	go grpcserver.InitialGRPC(grpcListener)
	var c = make(chan bool)
	<-c
}
