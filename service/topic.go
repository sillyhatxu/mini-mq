package service

import (
	"github.com/sillyhatxu/gocache-client"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/model"
	"github.com/sillyhatxu/mini-mq/utils/cache"
)

func CreateTopic(topic string) error {
	err := dbclient.InsertTopic(topic)
	if err != nil {
		return err
	}
	return dbclient.CreateTopicDataTable(topic)
}

func UpdateTopic(topic string, offset int64) error {
	value, found := cache.Client.Get(topic)
	if found {
		result := value.(*model.TopicDetail)
		result.Offset = offset
		cache.Client.Set(topic, &result, client.NoExpiration)
	}
	return dbclient.UpdateTopic(topic, offset)
}

func FindTopic(topic string) (*model.TopicDetail, error) {
	value, found := cache.Client.Get(topic)
	if found {
		result := value.(*model.TopicDetail)
		return result, nil
	}
	td, err := dbclient.FindByTopic(topic)
	if err != nil {
		return nil, err
	}
	if td != nil {
		cache.Client.Set(topic, td, client.NoExpiration)
	}
	return td, nil
}

func InsertTopicData(topic string, offset int64, body []byte) error {
	return dbclient.InsertTopicDataTransaction(topic, offset, body)
	//err := dbclient.InsertTopicData(topic, offset, body)
	//if err != nil {
	//	return err
	//}
	//return dbclient.UpdateTopic(topic, offset)
}
