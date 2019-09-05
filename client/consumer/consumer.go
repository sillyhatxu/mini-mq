package consumer

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/client"
	"github.com/sirupsen/logrus"
	"time"
)

type Delivery struct {
	TopicName      string
	TopicGroup     string
	LatestOffset   int64
	TopicDataArray []DeliveryData
}

type DeliveryData struct {
	TopicName  string
	TopicGroup string
	Offset     int64
	Body       []byte
}

type ConsumerConfig struct {
	Hearbeat *time.Duration
	NoWait   bool
	AutoAck  bool
}

type ConsumerClient struct {
	TopicName    string
	TopicGroup   string
	Offset       int64
	ConsumeCount int32
	Client       *client.Client
	Config       *ConsumerConfig
}

func (cc ConsumerClient) Validate() error {
	//TODO Validate
	return nil
}

type ConsumerInterface interface {
	MessageDelivery(msg Delivery)
}

func (cc ConsumerClient) Consume(ci ConsumerInterface) error {
	if cc.Client == nil {
		return fmt.Errorf("mq client is nil; ConsumerClient : %#v", cc)
	}
	msgs, err := cc.messageDelivery()
	if err != nil {
		logrus.Errorf("message delivery error; Error : %v", err)
		return nil
	}
	go func() {
		for delivery := range msgs {
			logrus.Infof("delivery : %#v", delivery)
			ci.MessageDelivery(delivery)
		}
	}()
	logrus.Info("waiting for messages.")
	forever := make(chan bool)
	<-forever
	logrus.Warningf("consumer exit; ConsumerClient : %#v", cc)
	return nil
}

func (cc ConsumerClient) messageDelivery() (<-chan Delivery, error) {
	time.Sleep(*cc.Config.Hearbeat)
	deliveryChannel := make(chan Delivery)
	body, err := cc.Client.GetTopicData(cc.TopicName, cc.TopicGroup, cc.Offset, cc.ConsumeCount)
	if err != nil {
		return nil, err
	}
	var resultArray []DeliveryData
	for _, td := range body.TopicDataArray {
		resultArray = append(resultArray, DeliveryData{
			TopicName:  td.TopicName,
			TopicGroup: td.TopicGroup,
			Offset:     td.Offset,
			Body:       td.Body,
		})
	}
	deliveryChannel <- *&Delivery{
		TopicName:      body.TopicName,
		TopicGroup:     body.TopicGroup,
		LatestOffset:   body.LatestOffset,
		TopicDataArray: resultArray,
	}
	return (<-chan Delivery)(deliveryChannel), nil
}
