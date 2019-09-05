package client

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
	Client := &Client{Address: Address, Timeout: 60 * time.Second}
	group, err := Client.Consume(TopicName, TopicGroup, Offset, ConsumeCount)
	assert.Nil(t, err)
	logrus.Infof("TopicGroup [%v] TopicName [%v] LatestOffset [%v]", group.TopicGroup, group.TopicName, group.LatestOffset)
	for i, td := range group.TopicDataArray {
		logrus.Infof("%d ;TopicGroup [%v] TopicName [%v] Offset [%v] Body : %v", i, td.TopicGroup, td.TopicName, td.Offset, string(td.Body))
	}
	err = Client.Commit(group.TopicName, group.TopicGroup, group.LatestOffset)
	assert.Nil(t, err)
}
