package dbclient

import (
	"database/sql"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sillyhatxu/mini-mq/model"
)

const (
	topicDataTable          = "topic_data_%s"
	createTopicDataTableSQL = `
CREATE TABLE IF NOT EXISTS %s
(
  id           INTEGER PRIMARY KEY AUTOINCREMENT,
  topic_name   TEXT NOT NULL,
  offset       INTEGER  DEFAULT 0,
  body         BLOB     DEFAULT 0,
  created_time datetime default current_timestamp
);
CREATE INDEX idx_%s_topic ON %s (topic_name);
	`
)

func getTopicDataTableName(topicName string) string {
	return fmt.Sprintf(topicDataTable, topicName)
}

func CreateTopicDataTable(topicName string) error {
	table := getTopicDataTableName(topicName)
	return Client.ExecDDL(fmt.Sprintf(createTopicDataTableSQL, table, table, table))
}

const dropTopicDataTableSQL = `
drop table %s
`

const deleteTopicGroupSQL = `
delete from topic_group where topic_name = ? and topic_group = ?
`

const deleteTopicDetailSQL = `
delete from topic_detail where topic_name = ?
`

func DeleteTopic(topicName string, topicGroupArray []model.TopicGroup) error {
	return Client.Transaction(func(tx *sql.Tx) error {
		err := Client.ExecDDL(fmt.Sprintf(dropTopicDataTableSQL, getTopicDataTableName(topicName)))
		if err != nil {
			return err
		}
		for _, tg := range topicGroupArray {
			stm, err := tx.Prepare(deleteTopicGroupSQL)
			if err != nil {
				return err
			}
			defer stm.Close()
			_, err = stm.Exec(tg.TopicName, tg.TopicGroup)
			if err != nil {
				return err
			}
		}
		stm, err := tx.Prepare(deleteTopicDetailSQL)
		if err != nil {
			return err
		}
		defer stm.Close()
		_, err = stm.Exec(topicName)
		if err != nil {
			return err
		}
		return nil
	})
}

const insertTopicData = `
insert into %s (topic_name,offset,body) values (?, ? , ?)
`

func InsertTopicData(topicName string, offset int64, body []byte) error {
	_, err := Client.Insert(fmt.Sprintf(insertTopicData, getTopicDataTableName(topicName)), topicName, offset, body)
	return err
}

func InsertTopicDataTransaction(topicName string, offset int64, body []byte) error {
	err := Client.Transaction(func(tx *sql.Tx) error {
		stm, err := tx.Prepare(fmt.Sprintf(insertTopicData, getTopicDataTableName(topicName)))
		if err != nil {
			return err
		}
		_, err = stm.Exec(topicName, offset, body)
		if err != nil {
			return err
		}
		err = stm.Close()
		if err != nil {
			return err
		}
		stm, err = tx.Prepare(updateTopicDetail)
		if err != nil {
			return err
		}
		_, err = stm.Exec(offset, topicName)
		if err != nil {
			return err
		}
		err = stm.Close()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

const insertTopicDetail = `
insert into topic_detail (topic_name) values (?)
`

func InsertTopicDetail(topicName string) error {
	_, err := Client.Insert(insertTopicDetail, topicName)
	return err
}

const updateTopicDetail = `
update topic_detail set offset = ?,last_modified_time = datetime('now') where topic_name = ?
`

func UpdateTopicDetail(topicName string, offset int64) error {
	_, err := Client.Update(updateTopicDetail, offset, topicName)
	return err
}

const findByTopicDetailList = `
select * from topic_detail
`

func FindByTopicDetailList() ([]model.TopicDetail, error) {
	var tdArray []model.TopicDetail
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02T15:04:05Z"),
		WeaklyTypedInput: true,
		Result:           &tdArray,
	}
	err := Client.FindListByConfig(findByTopicDetailList, config)
	if err != nil {
		return nil, err
	}
	if tdArray == nil {
		tdArray = make([]model.TopicDetail, 0)
	}
	return tdArray, err
}

const findByTopicDetail = `
select * from topic_detail where topic_name = ?
`

func FindByTopicDetail(topicName string) (*model.TopicDetail, error) {
	var td *model.TopicDetail
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02T15:04:05Z"),
		WeaklyTypedInput: true,
		Result:           &td,
	}
	err := Client.FindFirstByConfig(findByTopicDetail, config, topicName)
	return td, err
}

const findTopicGroupList = `
select * from topic_group limit ?,?
`

func FindTopicGroupList(offset int64, limit int) ([]model.TopicGroup, error) {
	var array []model.TopicGroup
	err := Client.FindList(findTopicGroupList, &array, offset, limit)
	if err != nil {
		return nil, err
	}
	if array == nil {
		return make([]model.TopicGroup, 0), nil
	}
	return array, nil
}

const findByTopicGroup = `
select * from topic_group where topic_name = ? and topic_group = ?
`

func FindByTopicGroup(topicName string, topicGroup string) (*model.TopicGroup, error) {
	var tg *model.TopicGroup
	err := Client.FindFirst(findByTopicGroup, &tg, topicName, topicGroup)
	return tg, err
}

const findByTopicGroupByTopicName = `
select * from topic_group where topic_name = ?
`

func FindByTopicGroupByTopicName(topicName string) ([]model.TopicGroup, error) {
	var array []model.TopicGroup
	err := Client.FindList(findByTopicGroupByTopicName, &array, topicName)
	if array == nil {
		array = make([]model.TopicGroup, 0)
	}
	return array, err
}

const insertTopicGroup = `
insert into topic_group (topic_name,topic_group,offset) values (?, ?, ?)
`

func InsertTopicGroup(topicName string, topicGroup string, offset int64) error {
	_, err := Client.Insert(insertTopicGroup, topicName, topicGroup, offset)
	return err
}

const updateTopicGroup = `
update topic_group set offset = ? where topic_name = ? and topic_group = ?
`

func UpdateTopicGroup(topicName string, topicGroup string, offset int64) error {
	_, err := Client.Update(updateTopicGroup, offset, topicName, topicGroup)
	return err
}

const findTopicData = `
select * from %s limit ?,?
`

func FindTopicData(topicName string, offset int64, consumeCount int) ([]model.TopicData, error) {
	var array []model.TopicData
	err := Client.FindList(fmt.Sprintf(findTopicData, getTopicDataTableName(topicName)), &array, offset, consumeCount)
	if err != nil {
		return nil, err
	}
	if array == nil {
		return make([]model.TopicData, 0), nil
	}
	return array, nil
}

const findTopicDataCount = `
select count(1) from %s
`

func FindTopicDataCount(topicName string) (int64, error) {
	return Client.Count(fmt.Sprintf(findTopicDataCount, getTopicDataTableName(topicName)))
}

func InsertTopicGroupTransaction(topic string, offset int64, body []byte) error {
	err := Client.Transaction(func(tx *sql.Tx) error {
		stm, err := tx.Prepare(fmt.Sprintf(insertTopicData, getTopicDataTableName(topic)))
		if err != nil {
			return err
		}
		_, err = stm.Exec(topic, offset, body)
		if err != nil {
			return err
		}
		err = stm.Close()
		if err != nil {
			return err
		}
		stm, err = tx.Prepare(updateTopicDetail)
		if err != nil {
			return err
		}
		_, err = stm.Exec(offset, topic)
		if err != nil {
			return err
		}
		err = stm.Close()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
