package main

import (
	"github.com/sillyhatxu/mini-mq/api"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"net"
)

func main() {
	dbclient.InitialDBClient("./basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	go api.InitialAPI(listener)
	//go grpc.InitialGRPC(listener)
	var c = make(chan bool)
	<-c
}
