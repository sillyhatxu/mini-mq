package producer

import (
	"github.com/sillyhatxu/mini-mq/client"
)

type ProducerClient struct {
	TopicName string
	Body      []byte
	Client    *client.Client
}

func (pc ProducerClient) Validate() error {
	//TODO Validate
	return nil
}

func (pc ProducerClient) Produce() error {
	err := pc.Validate()
	if err != nil {
		return err
	}
	return pc.Client.Produce(pc.TopicName, pc.Body)
}
