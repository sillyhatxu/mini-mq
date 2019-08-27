/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpctest/grpcproto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *grpcproto.HelloRequest) (*grpcproto.HelloResponse, error) {
	log.Printf("Received: %#v", in)
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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcproto.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
