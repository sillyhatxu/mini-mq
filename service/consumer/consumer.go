package consumer

import (
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/model"
	"github.com/sillyhatxu/mini-mq/service/topic"
)

type ConsumeGroup struct {
	TopicName    string
	TopicGroup   string
	Offset       int64
	ConsumeCount int
}

func NewConsumeGroup(TopicName string, TopicGroup string, Offset int64, ConsumeCount int) *ConsumeGroup {
	return &ConsumeGroup{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		Offset:       Offset,
		ConsumeCount: ConsumeCount,
	}
}

//type Consumer struct {
//	Topic        string
//	Group        string
//	Offset       *int64
//	ConsumeCount int
//	Timeout      time.Duration
//}

//chanel
func (cg *ConsumeGroup) Consume() ([]model.TopicData, error) {
	topicGroup, err := topic.FindTopicGroup(cg.TopicName, cg.TopicGroup, cg.Offset)
	if err != nil {
		return nil, err
	}
	topicDataArray, err := dbclient.FindTopicData(cg.TopicName, topicGroup.Offset, cg.ConsumeCount)
	if err != nil {
		return nil, err
	}
	cg.Offset = topicGroup.Offset + int64(len(topicDataArray))
	return topicDataArray, nil
}

func (cg *ConsumeGroup) Commit() error {
	return topic.UpdateTopicGroup(cg.TopicName, cg.TopicGroup, cg.Offset)
}
