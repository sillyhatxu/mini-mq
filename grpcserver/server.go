package grpcserver

import (
	"context"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/service/consumer"
	"github.com/sillyhatxu/mini-mq/service/producer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	CodeSuccess = "SUCCESS"
	CodeError   = "Error"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("---------- initial grpc server ----------")
	server := grpc.NewServer()
	RegisterProducerServiceServer(server, &produce{})
	RegisterConsumerServiceServer(server, &consume{})
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type produce struct{}

func (p *produce) Produce(ctx context.Context, in *ProduceRequest) (*ProduceResponse, error) {
	logrus.Infof("Received: %#v", in)
	err := producer.Produce(in.TopicName, in.Body)
	if err != nil {
		return &ProduceResponse{
			Code:    CodeError,
			Message: err.Error(),
		}, nil
	}
	return &ProduceResponse{
		Code:    CodeSuccess,
		Message: CodeSuccess,
	}, nil
}

type consume struct{}

func (p *consume) Consume(ctx context.Context, in *ConsumeRequest) (*ConsumeResponse, error) {
	logrus.Infof("Received: %#v", in)
	cg := consumer.NewConsumeGroup(in.TopicName, in.TopicGroup, in.Offset, int(in.ConsumeCount))
	array, err := cg.Consume()
	if err != nil {
		return &ConsumeResponse{
			Code:    CodeError,
			Message: err.Error(),
		}, nil
	}
	return &ConsumeResponse{
		Code:    CodeSuccess,
		Body:    toResponseBody(cg, array),
		Message: CodeSuccess,
	}, nil
}

func toResponseBody(consumeGroup *consumer.ConsumeGroup, array []dto.TopicData) *ConsumeResponse_Body {
	var resultArray []*ConsumeResponse_TopicData
	for _, td := range array {
		resultArray = append(resultArray, &ConsumeResponse_TopicData{
			TopicName:  td.TopicName,
			TopicGroup: td.TopicGroup,
			Offset:     td.Offset,
			Body:       td.Body,
		})
	}
	return &ConsumeResponse_Body{
		TopicName:      consumeGroup.TopicName,
		TopicGroup:     consumeGroup.TopicGroup,
		LatestOffset:   consumeGroup.Offset,
		TopicDataArray: resultArray,
	}
}

func (p *consume) Commit(ctx context.Context, in *CommitRequest) (*CommitResponse, error) {
	logrus.Infof("Received: %#v", in)
	cg := consumer.NewConsumeGroup(in.TopicName, in.TopicGroup, in.LatestOffset, 0)
	err := cg.Commit()
	if err != nil {
		return &CommitResponse{
			Code:    CodeError,
			Message: err.Error(),
		}, nil
	}
	return &CommitResponse{
		Code:    CodeSuccess,
		Message: CodeSuccess,
	}, nil
}
