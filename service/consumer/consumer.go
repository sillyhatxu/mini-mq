package consumer

import (
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/dto"
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

func (cg *ConsumeGroup) Consume() ([]dto.TopicData, error) {
	topicGroup, err := topic.FindTopicGroup(cg.TopicName, cg.TopicGroup, cg.Offset)
	if err != nil {
		return nil, err
	}
	topicDataArray, err := dbclient.FindTopicData(cg.TopicName, topicGroup.Offset, cg.ConsumeCount)
	if err != nil {
		return nil, err
	}
	cg.Offset = topicGroup.Offset + int64(len(topicDataArray))
	return toDTO(topicDataArray, cg.TopicGroup), nil
}

func toDTO(array []model.TopicData, topicGroup string) []dto.TopicData {
	var resultArray []dto.TopicData
	for _, td := range array {
		resultArray = append(resultArray, dto.TopicData{
			TopicName:  td.TopicName,
			TopicGroup: topicGroup,
			Offset:     td.Offset,
			Body:       td.Body,
		})
	}
	return resultArray
}

func (cg *ConsumeGroup) Commit() error {
	return topic.UpdateTopicGroup(cg.TopicName, cg.TopicGroup, cg.Offset)
}
