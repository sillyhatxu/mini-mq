package server

import (
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/mini-mq/grpc/consumer"
	"github.com/sillyhatxu/mini-mq/grpc/producer"
	"github.com/sillyhatxu/mini-mq/grpc/register"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("---------- initial grpc server ----------")
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus(consul.DefaultHealthCheckGRPCServerName, hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(server, healthServer)

	producer.RegisterProducerServiceServer(server, &producer.Producer{})
	consumer.RegisterConsumerServiceServer(server, &consumer.Consumer{})
	register.RegisterRegisterServiceServer(server, &register.Register{})
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
