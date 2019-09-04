package consumer

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	Address      = "localhost:8082"
	TopicName    = "test_topic"
	TopicGroup   = "test-1"
	Offset       = 0
	ConsumeCount = 5
)

func TestClient_Consume(t *testing.T) {
	Client := &Client{Address: Address, Timeout: 10 * time.Second}
	array, err := Client.Consume(Group{
		TopicName:    TopicName,
		TopicGroup:   TopicGroup,
		Offset:       Offset,
		ConsumeCount: ConsumeCount,
	})
	assert.Nil(t, err)
	for _, td := range array {
		logrus.Infof("td : %#v; body : %v", td, string(td.Body))
	}
}
