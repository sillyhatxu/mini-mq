package client

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

func (c *Client) getConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(c.Address, grpc.WithInsecure())
}

func (c Client) GetTopicData(TopicName string, TopicGroup string, Offset int64, ConsumeCount int32) (*grpcserver.ConsumeResponse_Body, error) {
	conn, err := c.getConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	consumerClient := grpcserver.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Consume(ctx, &grpcserver.ConsumeRequest{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		Offset:       Offset,
		ConsumeCount: ConsumeCount,
	})
	if err != nil {
		return nil, err
	}
	if response.Code != grpcserver.CodeSuccess {
		return nil, fmt.Errorf("consumer error; %v", response.Message)
	}
	return response.Body, nil
}

//func (c Client) Consume(TopicName string, TopicGroup string, Offset int64, ConsumeCount int) (*consumer.Group, error) {
//	body, err := c.getTopicData(TopicName, TopicGroup, Offset, int32(ConsumeCount))
//	if err != nil {
//		return nil, err
//	}
//	var resultArray []consumer.TopicData
//	for _, td := range body.TopicDataArray {
//		resultArray = append(resultArray, consumer.TopicData{
//			TopicName:  td.TopicName,
//			TopicGroup: td.TopicGroup,
//			Offset:     td.Offset,
//			Body:       td.Body,
//		})
//	}
//	return &consumer.Group{
//		TopicName:      body.TopicName,
//		TopicGroup:     body.TopicGroup,
//		LatestOffset:   body.LatestOffset,
//		TopicDataArray: resultArray,
//	}, nil
//}

func (c Client) Commit(TopicName string, TopicGroup string, LatestOffset int64) error {
	conn, err := c.getConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	consumerClient := grpcserver.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Commit(ctx, &grpcserver.CommitRequest{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		LatestOffset: LatestOffset,
	})
	if err != nil {
		return err
	}
	if response.Code != grpcserver.CodeSuccess {
		return fmt.Errorf("consumer commit error; %v", response.Message)
	}
	return nil
}

func (c Client) Produce(topicName string, body []byte) error {
	conn, err := c.getConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	producerClient := grpcserver.NewProducerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := producerClient.Produce(ctx, &grpcserver.ProduceRequest{
		TopicName: topicName,
		Body:      body,
	})
	if err != nil {
		return err
	}
	if response.Code != grpcserver.CodeSuccess {
		return fmt.Errorf("produce error; %v", response.Message)
	}
	return nil
}
