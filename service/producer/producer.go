package producer

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/service/topic"
	"sync"
)

func Produce(topicName string, body []byte) error {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicDetail, err := topic.FindTopic(topicName)
	if err != nil {
		return err
	}
	if topicDetail == nil {
		return fmt.Errorf("topicName[%s] does not exist, so you need to create topic first", topicName)
	}
	topicDetail.Offset++
	err = topic.InsertTopicData(topicDetail.TopicName, topicDetail.Offset, body)
	if err != nil {
		return err
	}
	return nil
}
