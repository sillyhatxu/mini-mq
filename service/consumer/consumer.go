package consumer

import (
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/model"
	"github.com/sillyhatxu/mini-mq/service/topic"
)

type Consumer struct {
	TopicName    string
	TopicGroup   string
	Offset       int64
	ConsumeCount int
}

//
//type Consumer struct {
//	Topic        string
//	Group        string
//	Offset       *int64
//	ConsumeCount int
//	Timeout      time.Duration
//}

//chanel
func (c *Consumer) Consume() ([]model.TopicData, error) {
	topicGroup, err := topic.FindTopicGroup(c.TopicName, c.TopicGroup, c.Offset)
	if err != nil {
		return nil, err
	}
	topicDataArray, err := dbclient.FindTopicData(c.TopicName, topicGroup.Offset, c.ConsumeCount)
	if err != nil {
		return nil, err
	}
	c.Offset = topicGroup.Offset + int64(len(topicDataArray))
	return topicDataArray, nil
}

func (c *Consumer) Commit() error {
	return topic.UpdateTopicGroup(c.TopicName, c.TopicGroup, c.Offset)
}
