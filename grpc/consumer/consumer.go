package consumer

import (
	"context"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/grpc/constants"
	"github.com/sillyhatxu/mini-mq/service/consumer"
	"github.com/sirupsen/logrus"
)

type Consumer struct{}

func (p *Consumer) Consume(ctx context.Context, in *ConsumeRequest) (*ConsumeResponse, error) {
	logrus.Infof("Received: %#v", in)
	cg := consumer.NewConsumeGroup(in.TopicName, in.TopicGroup, in.Offset, int(in.ConsumeCount))
	array, err := cg.Consume()
	if err != nil {
		return &ConsumeResponse{
			Code:    constants.CodeError,
			Message: err.Error(),
		}, nil
	}
	return &ConsumeResponse{
		Code:    constants.CodeSuccess,
		Body:    toResponseBody(cg, array),
		Message: constants.CodeSuccess,
	}, nil
}

func toResponseBody(consumeGroup *consumer.ConsumeGroup, array []dto.TopicData) *ConsumeResponse_Body {
	var resultArray []*ConsumeResponse_TopicData
	for _, td := range array {
		resultArray = append(resultArray, &ConsumeResponse_TopicData{
			TopicName:  td.TopicName,
			TopicGroup: td.TopicGroup,
			Offset:     td.Offset,
			Body:       td.Body,
		})
	}
	return &ConsumeResponse_Body{
		TopicName:      consumeGroup.TopicName,
		TopicGroup:     consumeGroup.TopicGroup,
		LatestOffset:   consumeGroup.Offset,
		TopicDataArray: resultArray,
	}
}

func (p *Consumer) Commit(ctx context.Context, in *CommitRequest) (*CommitResponse, error) {
	logrus.Infof("Received: %#v", in)
	cg := consumer.NewConsumeGroup(in.TopicName, in.TopicGroup, in.LatestOffset, 0)
	err := cg.Commit()
	if err != nil {
		return &CommitResponse{
			Code:    constants.CodeError,
			Message: err.Error(),
		}, nil
	}
	return &CommitResponse{
		Code:    constants.CodeSuccess,
		Message: constants.CodeSuccess,
	}, nil
}
