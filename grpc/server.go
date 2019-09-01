package grpc

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpctest/grpcproto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("---------- initial grpc server ----------")
	grpcServer := grpc.NewServer()
	grpcproto.RegisterUserServiceServer(grpcServer, &server{})
	err := grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *grpcproto.HelloRequest) (*grpcproto.HelloResponse, error) {
	logrus.Infof("Received: %#v", in)
	return &grpcproto.HelloResponse{
		Code: "SUCCESS",
		Body: &grpcproto.HelloResponse_User{
			UserId:     in.UserId,
			UserName:   "UserName",
			Age:        23,
			IsDelete:   true,
			CreateTime: time.Now().UnixNano() / int64(time.Millisecond),
			Msg:        []byte("This is test msg"),
			AddressList: []*grpcproto.HelloResponse_Address{
				{
					Id:       1,
					Address:  "address 111111",
					ZipCode:  "111111",
					Photo:    "11111111",
					IsDelete: false,
				},
				{
					Id:       2,
					Address:  "address 222222",
					ZipCode:  "222222",
					Photo:    "22222222",
					IsDelete: false,
				},
				{
					Id:       3,
					Address:  "address 333333",
					ZipCode:  "333333",
					Photo:    "33333333",
					IsDelete: false,
				},
			},
		},
		Message: "Success",
	}, nil
}
