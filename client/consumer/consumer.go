package consumer

import (
	"context"
	"fmt"
	"github.com/sillyhatxu/mini-mq/grpcserver"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	Address string
	Timeout time.Duration
}

type Group struct {
	TopicName    string
	TopicGroup   string
	Offset       int64
	ConsumeCount int
}

type GroupCommit struct {
	TopicName    string
	TopicGroup   string
	LatestOffset int64
}

type TopicData struct {
	TopicName  string
	TopicGroup string
	Offset     int64
	Body       []byte
}

func (c Client) Consume(group Group) ([]TopicData, error) {
	array, err := c.getTopicData(group)
	if err != nil {
		return nil, err
	}
	var resultArray []TopicData
	for _, td := range array {
		resultArray = append(resultArray, TopicData{
			TopicName:  td.TopicName,
			TopicGroup: td.TopicGroup,
			Offset:     td.Offset,
			Body:       td.Body,
		})
	}
	return resultArray, nil
}

func (c Client) getTopicData(group Group) ([]*grpcserver.ConsumeResponse_Body, error) {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	consumerClient := grpcserver.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Consume(ctx, &grpcserver.ConsumeRequest{
		TopicName:    group.TopicName,
		TopicGroup:   group.TopicGroup,
		Offset:       group.Offset,
		ConsumeCount: int32(group.ConsumeCount),
	})
	if err != nil {
		return nil, err
	}
	if response.Code != grpcserver.CodeSuccess {
		return nil, fmt.Errorf("consumer error; %v", response.Message)
	}
	return response.Body, nil
}

func (c Client) Commit(group GroupCommit) error {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	consumerClient := grpcserver.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Commit(ctx, &grpcserver.CommitRequest{
		TopicName:    group.TopicName,
		TopicGroup:   group.TopicGroup,
		LatestOffset: group.LatestOffset,
	})
	if err != nil {
		return err
	}
	if response.Code != grpcserver.CodeSuccess {
		return fmt.Errorf("consumer commit error; %v", response.Message)
	}
	return nil
}
