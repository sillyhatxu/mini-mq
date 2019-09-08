package client

import (
	"context"
	"fmt"
	"github.com/sillyhatxu/mini-mq/grpc/constants"
	"github.com/sillyhatxu/mini-mq/grpc/consumer"
	"github.com/sillyhatxu/mini-mq/grpc/producer"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	Address string
	Timeout time.Duration
}

func (c *Client) GetConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(c.Address, grpc.WithInsecure())
}

func (c Client) GetTopicData(TopicName string, TopicGroup string, Offset int64, ConsumeCount int32) (*consumer.ConsumeResponse_Body, error) {
	conn, err := c.GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	consumerClient := consumer.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Consume(ctx, &consumer.ConsumeRequest{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		Offset:       Offset,
		ConsumeCount: ConsumeCount,
	})
	if err != nil {
		return nil, err
	}
	if response.Code != constants.CodeSuccess {
		return nil, fmt.Errorf("consumer error; %v", response.Message)
	}
	return response.Body, nil
}

func (c Client) Commit(topicName string, topicGroup string, latestOffset int64) error {
	conn, err := c.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	consumerClient := consumer.NewConsumerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := consumerClient.Commit(ctx, &consumer.CommitRequest{
		TopicName:    topicName,
		TopicGroup:   topicGroup,
		LatestOffset: latestOffset,
	})
	if err != nil {
		return err
	}
	if response.Code != constants.CodeSuccess {
		return fmt.Errorf("consumer commit error; %v", response.Message)
	}
	return nil
}

func (c Client) Produce(topicName string, body []byte) error {
	conn, err := c.GetConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	producerClient := producer.NewProducerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	response, err := producerClient.Produce(ctx, &producer.ProduceRequest{
		TopicName: topicName,
		Body:      body,
	})
	if err != nil {
		return err
	}
	if response.Code != constants.CodeSuccess {
		return fmt.Errorf("produce error; %v", response.Message)
	}
	return nil
}
