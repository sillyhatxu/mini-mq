package producer

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpc/constants"
	"github.com/sillyhatxu/mini-mq/service/producer"
	"github.com/sirupsen/logrus"
)

type Producer struct{}

func (p *Producer) Produce(ctx context.Context, in *ProduceRequest) (*ProduceResponse, error) {
	logrus.Infof("Received: %#v", in)
	err := producer.Produce(in.TopicName, in.Body)
	if err != nil {
		return &ProduceResponse{
			Code:    constants.CodeError,
			Message: err.Error(),
		}, nil
	}
	return &ProduceResponse{
		Code:    constants.CodeSuccess,
		Message: constants.CodeSuccess,
	}, nil
}
