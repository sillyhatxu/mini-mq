package register

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpc/constants"
	"github.com/sillyhatxu/mini-mq/service/register"
	"github.com/sirupsen/logrus"
)

type Register struct{}

func (p *Register) Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	logrus.Infof("Received: %#v", in)
	//TODO return error check
	//err := register.Register(in.GetTopicName(), in.GetTopicGroup(), in.Address)
	err := register.Register(in.GetTopicName(), in.GetTopicGroup(), in.Address)
	if err != nil {
		return &RegisterResponse{
			Code:    constants.CodeError,
			Message: err.Error(),
		}, nil
	}
	return &RegisterResponse{
		Code:    constants.CodeSuccess,
		Message: constants.CodeSuccess,
	}, nil
}
