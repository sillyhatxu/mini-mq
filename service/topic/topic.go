package topic

import (
	"fmt"
	"github.com/sillyhatxu/gocache-client"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/model"
	"github.com/sillyhatxu/mini-mq/utils/cache"
)

const (
	topicDetailKey = "TOPIC_DETAIL_%s"
	topicGroupKey  = "TOPIC_GROUP_%s_%s"
)

func getTopicDetailKey(topicName string) string {
	return fmt.Sprintf(topicDetailKey, topicName)
}

func getTopicGroupKey(topicName string, topicGroup string) string {
	return fmt.Sprintf(topicGroupKey, topicName, topicGroup)
}

func CreateTopic(topicName string) error {
	err := dbclient.InsertTopicDetail(topicName)
	if err != nil {
		return err
	}
	return dbclient.CreateTopicDataTable(topicName)
}

func UpdateTopic(topicName string, offset int64) error {
	value, found := cache.Client.Get(topicName)
	if found {
		result := value.(*model.TopicDetail)
		result.Offset = offset
		cache.Client.Set(topicName, &result, client.NoExpiration)
	}
	return dbclient.UpdateTopicDetail(topicName, offset)
}

func FindTopic(topicName string) (*model.TopicDetail, error) {
	value, found := cache.Client.Get(getTopicDetailKey(topicName))
	if found {
		result := value.(*model.TopicDetail)
		return result, nil
	}
	td, err := dbclient.FindByTopicDetail(topicName)
	if err != nil {
		return nil, err
	}
	if td != nil {
		cache.Client.Set(getTopicDetailKey(topicName), td, client.NoExpiration)
	}
	return td, nil
}

func FindTopicGroup(topicName string, topicGroup string, offset int64) (tg *model.TopicGroup, err error) {
	value, found := cache.Client.Get(getTopicGroupKey(topicName, topicGroup))
	if found {
		result := value.(*model.TopicGroup)
		return result, nil
	}
	tg, err = dbclient.FindByTopicGroup(topicName, topicGroup)
	if err != nil {
		return nil, err
	}
	if tg == nil {
		err := dbclient.InsertTopicGroup(topicName, topicGroup, offset)
		if err != nil {
			return nil, err
		}
		tg = &model.TopicGroup{
			TopicName:  topicName,
			TopicGroup: topicGroup,
			Offset:     offset,
		}
	}
	cache.Client.Set(getTopicGroupKey(topicName, topicGroup), tg, client.NoExpiration)
	return tg, nil
}

func InsertTopicData(topicName string, offset int64, body []byte) error {
	return dbclient.InsertTopicDataTransaction(topicName, offset, body)
	//err := dbclient.InsertTopicData(topic, offset, body)
	//if err != nil {
	//	return err
	//}
	//return dbclient.UpdateTopic(topic, offset)
}

func UpdateTopicGroup(topicName string, topicGroup string, offset int64) error {
	tg, err := FindTopicGroup(topicName, topicGroup, offset)
	if err != nil {
		return err
	}
	tg.Offset = offset
	err = dbclient.UpdateTopicGroup(topicName, topicGroup, offset)
	if err != nil {
		return err
	}
	cache.Client.Set(getTopicGroupKey(topicName, topicGroup), tg, client.NoExpiration)
	return nil
}
