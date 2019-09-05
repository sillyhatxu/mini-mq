package register

import (
	"context"
	"github.com/sillyhatxu/mini-mq/grpc/constants"
	"github.com/sirupsen/logrus"
)

type Register struct{}

func (p *Register) Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	logrus.Infof("Received: %#v", in)

	return &RegisterResponse{
		Code:    constants.CodeSuccess,
		Message: constants.CodeSuccess,
	}, nil
}
