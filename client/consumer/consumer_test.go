package consumer

import (
	"github.com/sillyhatxu/mini-mq/client"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	Address      = "localhost:8082"
	TopicName    = "test_topic"
	TopicGroup   = "test-3"
	Offset       = 0
	ConsumeCount = 5
)

type ConsumerTest struct{}

func (ct ConsumerTest) MessageDelivery(delivery Delivery) error {
	logrus.Infof("delivery { TopicName:%v, TopicGroup:%v, LatestOffset:%v }", delivery.TopicName, delivery.TopicGroup, delivery.LatestOffset)
	return nil
}

func TestClient_Consume(t *testing.T) {
	Client := &ConsumerClient{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		Offset:       Offset,
		ConsumeCount: ConsumeCount,
		Client: &client.Client{
			Address: Address,
			Timeout: 60 * time.Second,
		},
		Config: &ConsumerConfig{
			Hearbeat: 5 * time.Second,
			NoWait:   true,
			AutoAck:  true,
		},
	}
	err := Client.Consume(&ConsumerTest{})
	if err != nil {
		panic(err)
	}
}
