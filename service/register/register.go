package register

import (
	"github.com/sillyhatxu/mini-mq/service/consuming"
)

func Register(topicName string, topicGroup string, address string) error {
	rc := consuming.RegisterConsumer{
		TopicName:  topicName,
		TopicGroup: topicGroup,
		Address:    address,
	}
	return rc.Register()
}
