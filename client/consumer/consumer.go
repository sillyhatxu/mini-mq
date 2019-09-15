package consumer

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/client/client"
	"github.com/sirupsen/logrus"
	"time"
)

type ConsumerClient struct {
	client       *client.Client
	topicName    string
	topicGroup   string
	offset       int64
	consumeCount int32
	config       *ConsumerConfig
}

type ConsumerConfig struct {
	hearbeat time.Duration
	noWait   bool
	autoAck  bool
}

func NewConsumerClient(client *client.Client, topicName string, topicGroup string, offset int64, consumeCount int32, opts ...Option) *ConsumerClient {
	//default
	config := &ConsumerConfig{
		hearbeat: 5 * time.Second,
		noWait:   true,
		autoAck:  true,
	}
	for _, opt := range opts {
		opt(config)
	}
	return &ConsumerClient{
		client:       client,
		topicName:    topicName,
		topicGroup:   topicGroup,
		offset:       offset,
		consumeCount: consumeCount,
		config:       config,
	}
}

type Option func(*ConsumerConfig)

func Hearbeat(hearbeat time.Duration) Option {
	return func(c *ConsumerConfig) {
		c.hearbeat = hearbeat
	}
}

func NoWait(noWait bool) Option {
	return func(c *ConsumerConfig) {
		c.noWait = noWait
	}
}

func AutoAck(autoAck bool) Option {
	return func(c *ConsumerConfig) {
		c.autoAck = autoAck
	}
}

func (cc ConsumerClient) Validate() error {
	//TODO Validate
	return nil
}

type ConsumerInterface interface {
	MessageDelivery(delivery Delivery) error
}

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

func (cc ConsumerClient) Consume(ci ConsumerInterface) error {
	err := cc.Validate()
	if err != nil {
		return err
	}
	if cc.client == nil {
		return fmt.Errorf("mq client is nil; ConsumerClient : %#v", cc)
	}
	msgs, err := cc.messageDelivery()
	if err != nil {
		logrus.Errorf("message delivery error; Error : %v", err)
		return nil
	}
	go func() {
		for delivery := range msgs {
			//logrus.Infof("delivery : %#v", delivery)
			err := ci.MessageDelivery(delivery)
			if err != nil {
				logrus.Errorf("message delivery; Error : %v", err)
				continue
			}
			if cc.config.autoAck {
				err := cc.client.Commit(cc.topicName, cc.topicGroup, delivery.LatestOffset)
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
		for {
			time.Sleep(cc.config.hearbeat)
			delivery, err := cc.getDelivery()
			if err != nil {
				continue
			}
			if delivery.TopicDataArray == nil || len(delivery.TopicDataArray) == 0 {
				continue
			}
			deliveryChannel <- *delivery
		}
	}()
	return (<-chan Delivery)(deliveryChannel), nil
}

func (cc ConsumerClient) getDelivery() (*Delivery, error) {
	body, err := cc.client.GetTopicData(cc.topicName, cc.topicGroup, cc.offset, cc.consumeCount)
	if err != nil {
		logrus.Errorf("get topic data error; %v", err)
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
	return &Delivery{
		TopicName:      body.TopicName,
		TopicGroup:     body.TopicGroup,
		LatestOffset:   body.LatestOffset,
		TopicDataArray: resultArray,
	}, nil
}
