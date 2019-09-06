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
	MessageDelivery(msg Delivery) error
}

func (cc ConsumerClient) Consume(ci ConsumerInterface) error {
	err := cc.Validate()
	if err != nil {
		return err
	}
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
			err := ci.MessageDelivery(delivery)
			if err != nil {
				logrus.Errorf("message delivery; Error : %v", err)
				continue
			}
			if cc.Config.AutoAck {
				err := cc.Client.Commit(cc.TopicName, cc.TopicGroup, delivery.LatestOffset)
				if err != nil {
					logrus.Errorf("message delivery; Error : %v", err)
					continue
				}
			}
		}
	}()
	logrus.Info("waiting for messages.")
	forever := make(chan bool)
	<-forever
	logrus.Warningf("consumer exit; ConsumerClient : %#v", cc)
	return nil
}

func (cc ConsumerClient) messageDelivery() (<-chan Delivery, error) {
	deliveryChannel := make(chan Delivery)
	go func() {
		time.Sleep(*cc.Config.Hearbeat)
		body, err := cc.Client.GetTopicData(cc.TopicName, cc.TopicGroup, cc.Offset, cc.ConsumeCount)
		if err != nil {
			logrus.Errorf("get topic data error; %v", err)
			return
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
	}()
	return (<-chan Delivery)(deliveryChannel), nil
}
