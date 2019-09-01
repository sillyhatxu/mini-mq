package consumer

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

const path = "/Users/shikuanxu/go/src/github.com/sillyhatxu/mini-mq"

//const path = "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq"

func TestConsumer(t *testing.T) {
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	cache.Initial()
	c := Consumer{
		TopicName:    "test_topic",
		TopicGroup:   "test-1",
		Offset:       0,
		ConsumeCount: 10,
	}
	array, err := c.Consume()
	assert.Nil(t, err)
	for i, data := range array {
		logrus.Infof("%d ; TopicName [%v] Offset [%v] Body : %v", i, data.TopicName, data.Offset, string(data.Body))
	}
	err = c.Commit()
	assert.Nil(t, err)
}

func TestConsumerBatch(t *testing.T) {
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	cache.Initial()
	c := Consumer{
		TopicName:    "test_topic",
		TopicGroup:   "test-1",
		Offset:       0,
		ConsumeCount: 10,
	}
	for i := 0; i < 5; i++ {
		array, err := c.Consume()
		assert.Nil(t, err)
		for i, data := range array {
			logrus.Infof("%d ; TopicName [%v] Offset [%v] Body : %v", i, data.TopicName, data.Offset, string(data.Body))
		}
		err = c.Commit()
		assert.Nil(t, err)
	}
}
