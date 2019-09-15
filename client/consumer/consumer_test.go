package consumer

import (
	"github.com/sillyhatxu/mini-mq/client/client"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

const (
	Address      = "localhost:8082"
	TopicName    = "test_topic"
	TopicGroup   = "test-1"
	Offset       = 0
	ConsumeCount = 5
)

type ConsumerTest struct{}

func (ct ConsumerTest) MessageDelivery(delivery Delivery) error {
	logrus.Infof("delivery { TopicName:%v, TopicGroup:%v, LatestOffset:%v , data length : %d}", delivery.TopicName, delivery.TopicGroup, delivery.LatestOffset, len(delivery.TopicDataArray))

	return nil
}

var Client = client.NewClient("localhost:8200", client.Timeout(30*time.Second))

func TestClient_Consume(t *testing.T) {
	consume := NewConsumerClient(Client, TopicName, TopicGroup, Offset, ConsumeCount)
	err := consume.Consume(&ConsumerTest{})
	if err != nil {
		panic(err)
	}
}
