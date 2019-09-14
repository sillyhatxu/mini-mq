package topic

import (
	"fmt"
	"github.com/sillyhatxu/gocache-client"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/model"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	topicDetailKey = "TOPIC_DETAIL"
	topicGroupKey  = "TOPIC_GROUP"
)

//TODO channel
func getTopicDetailMap() map[string]*model.TopicDetail {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	for {
		value, found := cache.Client.Get(topicDetailKey)
		if !found {
			logrus.Infof("initial topic detail map[%s]", topicDetailKey)
			topicDetailMap := make(map[string]*model.TopicDetail)
			cache.Client.Set(topicDetailKey, topicDetailMap, client.NoExpiration)
			continue
		}
		return value.(map[string]*model.TopicDetail)
	}
}

func setTopicDetailMap(topicName string, td *model.TopicDetail) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicDetailMap := getTopicDetailMap()
	topicDetailMap[topicName] = td
	cache.Client.Set(topicDetailKey, topicDetailMap, client.NoExpiration)
}

func deleteTopicDetailMap(topicName string) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicDetailMap := getTopicDetailMap()
	delete(topicDetailMap, topicName)
	cache.Client.Set(topicDetailKey, topicDetailMap, client.NoExpiration)
}

func FindTopicList() ([]model.TopicDetail, error) {
	return dbclient.FindByTopicDetailList()
}

func FindTopic(topicName string) (*model.TopicDetail, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicDetailMap := getTopicDetailMap()
	topicDetail, ok := topicDetailMap[topicName]
	if ok {
		return topicDetail, nil
	}
	td, err := dbclient.FindByTopicDetail(topicName)
	if err != nil {
		return nil, err
	}
	if td != nil {
		setTopicDetailMap(topicName, td)
	}
	return td, nil
}

func CreateTopic(topicName string) error {
	td, err := FindTopic(topicName)
	if err != nil {
		return err
	}
	if td != nil {
		return fmt.Errorf("[%s]topic name has been exist", topicName)
	}
	err = dbclient.InsertTopicDetail(topicName)
	if err != nil {
		return err
	}
	err = dbclient.CreateTopicDataTable(topicName)
	if err != nil {
		return err
	}
	setTopicDetailMap(topicName, &model.TopicDetail{TopicName: topicName, Offset: 0, CreatedTime: time.Now(), LastModifiedTime: time.Now()})
	return nil
}

func UpdateTopic(topicName string, offset int64) error {
	td, err := FindTopic(topicName)
	if err != nil {
		return err
	}
	if td == nil {
		return fmt.Errorf("topic[%s] does not exist", topicName)
	}
	err = dbclient.UpdateTopicDetail(topicName, offset)
	if err != nil {
		return err
	}
	td.Offset = offset
	setTopicDetailMap(topicName, td)
	return nil
}

func DeleteTopic(topicName string) error {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	td, err := FindTopic(topicName)
	if err != nil {
		return err
	}
	if td == nil {
		return fmt.Errorf("topic[%s] does not exist", topicName)
	}
	topicGroupArray, err := dbclient.FindByTopicGroupByTopicName(topicName)
	if err != nil {
		return err
	}
	for _, topicGroup := range topicGroupArray {
		deleteTopicGroupMap(&topicGroup)
	}
	deleteTopicDetailMap(topicName)
	return dbclient.DeleteTopic(topicName, topicGroupArray)
}

func FindTopicData(topicName string, offset int64, limit int) ([]dto.TopicDataDTO, int64, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	td, err := FindTopic(topicName)
	if err != nil {
		return nil, 0, err
	}
	if td == nil {
		return nil, 0, fmt.Errorf("topic[%s] does not exist", topicName)
	}
	count, err := dbclient.FindTopicDataCount(topicName)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return make([]dto.TopicDataDTO, 0), 0, nil
	}
	array, err := dbclient.FindTopicData(topicName, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return switchTopicDataToDTO(array), count, nil
}

func switchTopicDataToDTO(array []model.TopicData) []dto.TopicDataDTO {
	var resultArray []dto.TopicDataDTO
	for _, td := range array {
		resultArray = append(resultArray, dto.TopicDataDTO{
			TopicName: td.TopicName,
			Offset:    td.Offset,
			Body:      string(td.Body),
		})
	}
	return resultArray
}

func getTopicGroupMap() map[string]*model.TopicGroup {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	for {
		value, found := cache.Client.Get(topicGroupKey)
		if !found {
			logrus.Infof("initial topic group map[%s]", topicGroupKey)
			continue
		}
		return value.(map[string]*model.TopicGroup)
	}
}

//TODO channel
func getTopicGroupMapKey(topicName string, topicGroup string) string {
	return fmt.Sprintf("%s_%s", topicName, topicGroup)
}

func setTopicGroupMap(topicName string, topicGroup string, tg *model.TopicGroup) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicGroupMap := getTopicGroupMap()
	topicGroupMap[getTopicGroupMapKey(topicName, topicGroup)] = tg
	cache.Client.Set(topicGroupKey, topicGroupMap, client.NoExpiration)
}

func deleteTopicGroupMap(tg *model.TopicGroup) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	topicGroupMap := getTopicGroupMap()
	delete(topicGroupMap, getTopicGroupMapKey(tg.TopicName, tg.TopicGroup))
	cache.Client.Set(topicGroupKey, topicGroupMap, client.NoExpiration)
}

func FindTopicGroup(topicName string, topicGroup string, offset int64) (*model.TopicGroup, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	td, err := FindTopic(topicName)
	if err != nil {
		return nil, err
	}
	if td == nil {
		return nil, fmt.Errorf("topic[%s] does not exist", topicName)
	}
	topicGroupMap := getTopicGroupMap()
	tg, ok := topicGroupMap[getTopicGroupMapKey(topicName, topicGroup)]
	if ok {
		return tg, nil
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
	setTopicGroupMap(topicName, topicGroup, tg)
	return tg, nil
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
	setTopicGroupMap(topicName, topicGroup, tg)
	return nil
}

func InsertTopicData(topicName string, offset int64, body []byte) error {
	return dbclient.InsertTopicDataTransaction(topicName, offset, body)
	//err := dbclient.InsertTopicData(topic, offset, body)
	//if err != nil {
	//	return err
	//}
	//return dbclient.UpdateTopic(topic, offset)
}
