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
	cg := NewConsumeGroup("test_topic", "test-1", 0, 10)
	array, err := cg.Consume()
	assert.Nil(t, err)
	for i, data := range array {
		logrus.Infof("%d ; TopicName [%v] Offset [%v] Body : %v", i, data.TopicName, data.Offset, string(data.Body))
	}
	err = cg.Commit()
	assert.Nil(t, err)
}

func TestConsumerBatch(t *testing.T) {
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	cache.Initial()
	cg := NewConsumeGroup("test_topic", "test-3", 0, 10)
	for i := 0; i < 5; i++ {
		array, err := cg.Consume()
		assert.Nil(t, err)
		for i, data := range array {
			logrus.Infof("%d ; TopicName [%v] Offset [%v] Body : %v", i, data.TopicName, data.Offset, string(data.Body))
		}
		err = cg.Commit()
		assert.Nil(t, err)
	}
}
