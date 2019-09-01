package main

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpctest/grpcproto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:8082"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpcproto.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &grpcproto.HelloRequest{UserId: 123})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("FindUserByUserId : %#v", r.Message)
	log.Printf("FindUserByUserId response : %#v", r.Body)

	log.Printf("UserId : %v", r.Body.UserId)
	log.Printf("UserName : %v", r.Body.UserName)
	log.Printf("Age : %v", r.Body.Age)
	log.Printf("IsDelete : %v", r.Body.IsDelete)
	log.Printf("CreateTime : %v", r.Body.CreateTime)
	log.Printf("Msg : %v", string(r.Body.Msg))
	for i, address := range r.Body.AddressList {
		log.Printf("%d - %#v", i, address)
	}

}
