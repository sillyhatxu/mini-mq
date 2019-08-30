package service

import (
	"fmt"
)

func Produce(topic string, body []byte) error {
	//var mutex sync.Mutex
	//mutex.Lock()
	//defer mutex.Unlock()
	topicDetail, err := FindTopic(topic)
	if err != nil {
		return err
	}
	if topicDetail == nil {
		return fmt.Errorf("topic[%s] does not exist, so you need to create topic first", topic)
	}
	topicDetail.Offset++
	err = InsertTopicData(topicDetail.Topic, topicDetail.Offset, body)
	if err != nil {
		return err
	}
	return nil
}
