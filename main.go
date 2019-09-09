package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/mini-mq/api"
	"github.com/sillyhatxu/mini-mq/config"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/grpc/server"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"log"
	"net"
)

func init() {
	cfgFile := flag.String("c", "config-local.conf", "config file")
	flag.Parse()
	err := envconfig.ParseEnvironmentConfig(&config.Conf.EnvConfig)
	if err != nil {
		panic(err)
	}
	envconfig.ParseConfig(*cfgFile, func(content []byte) {
		err := toml.Unmarshal(content, &config.Conf)
		if err != nil {
			panic(fmt.Sprintf("unmarshal toml object error. %v", err))
		}
	})
	log.Printf("config.Conf : %#v", config.Conf)
	config.InitialLogConfig()
}

func main() {
	dbclient.InitialDBClient(config.Conf.DB.DataSourceName, config.Conf.DB.DDLPath)
	cache.Initial()
	apiListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.HttpPort))
	if err != nil {
		panic(err)
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.GRPCPort))
	if err != nil {
		panic(err)
	}
	go api.InitialAPI(apiListener)
	go server.InitialGRPC(grpcListener)
	var c = make(chan bool)
	<-c
}
