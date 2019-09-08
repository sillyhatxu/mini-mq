package consuming

import (
	"fmt"
	"github.com/sillyhatxu/gocache-client"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/service/topic"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"github.com/sirupsen/logrus"
)

const (
	consumingKey        = "CONSUMING"
	consumerGroupKey    = "%s_%s_%s"
	consumingStatusUp   = "UP"
	consumingStatusDown = "Down"
)

func getConsumerGroupKey(topicName string, topicGroup string) string {
	return fmt.Sprintf(consumerGroupKey, consumingKey, topicName, topicGroup)
}

type RegisterConsumer struct {
	TopicName  string
	TopicGroup string
	Address    string
}

type Consuming struct {
	Address string
	Status  string
}

func List() []dto.ConsumingDTO {
	var array []dto.ConsumingDTO
	consumingMap := GetConsumingMap()
	for key, value := range consumingMap {
		array = append(array, dto.ConsumingDTO{
			ConsumingKey: key,
			Address:      value.Address,
			Status:       value.Status,
		})
	}
	if array == nil {
		array = make([]dto.ConsumingDTO, 0)
	}
	return array
}

func GetConsumingMap() map[string]*Consuming {
	for {
		value, found := cache.Client.Get(consumingKey)
		if !found {
			logrus.Infof("initial consuming map[%s]", consumingKey)
			continue
		}
		return value.(map[string]*Consuming)
	}
}

func (rc RegisterConsumer) Register() error {
	td, err := topic.FindTopic(rc.TopicName)
	if err != nil {
		return nil
	}
	if td == nil {
		return fmt.Errorf("topic[%s] does not exist", rc.TopicName)
	}
	consumingMap := GetConsumingMap()
	consuming, ok := consumingMap[getConsumerGroupKey(rc.TopicName, rc.TopicGroup)]
	if ok && consuming.Status == consumingStatusUp {
		return fmt.Errorf("consumer has been exist. TopicName : %s; TopicGroup : %s", rc.TopicName, rc.TopicGroup)
	}
	consumingMap[getConsumerGroupKey(rc.TopicName, rc.TopicGroup)] = &Consuming{Status: consumingStatusDown, Address: rc.Address}
	rc.SetConsumingMap(consumingMap)
	return nil
}

func (rc RegisterConsumer) SetConsumingMap(consumingMap map[string]*Consuming) {
	cache.Client.Set(consumingKey, consumingMap, client.NoExpiration)
}
