package server

import (
	"github.com/sillyhatxu/mini-mq/grpc/consumer"
	"github.com/sillyhatxu/mini-mq/grpc/health"
	"github.com/sillyhatxu/mini-mq/grpc/producer"
	"github.com/sillyhatxu/mini-mq/grpc/register"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("---------- initial grpc server ----------")
	server := grpc.NewServer()
	producer.RegisterProducerServiceServer(server, &producer.Producer{})
	consumer.RegisterConsumerServiceServer(server, &consumer.Consumer{})
	register.RegisterRegisterServiceServer(server, &register.Register{})
	go health.Check()
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
