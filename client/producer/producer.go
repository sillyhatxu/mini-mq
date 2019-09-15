package producer

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/client/client"
)

type ProducerClient struct {
	topicName string
	client    *client.Client
}

func NewProducerClient(client *client.Client, topicName string) *ProducerClient {
	return &ProducerClient{
		client:    client,
		topicName: topicName,
	}
}

func (pc ProducerClient) Validate() error {
	if pc.client == nil {
		return fmt.Errorf("client is nil")
	} else if pc.topicName == "" {
		return fmt.Errorf("topicName is nil")
	}
	return nil
}

func (pc ProducerClient) Produce(body []byte) error {
	err := pc.Validate()
	if err != nil {
		return err
	}
	return pc.client.Produce(pc.topicName, body)
}
