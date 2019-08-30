package dbclient

import (
	"database/sql"
	"fmt"
	"github.com/sillyhatxu/mini-mq/model"
)

const (
	topicDataTable          = "topic_data_%s"
	createTopicDataTableSQL = `
CREATE TABLE IF NOT EXISTS %s
(
  id           INTEGER PRIMARY KEY AUTOINCREMENT,
  topic        TEXT NOT NULL,
  offset       INTEGER  DEFAULT 0,
  body         BLOB     DEFAULT 0,
  created_time datetime default current_timestamp
);
CREATE INDEX idx_%s_topic ON %s (topic);
	`
)

func getTopicTableName(topic string) string {
	return fmt.Sprintf(topicDataTable, topic)
}

func CreateTopicDataTable(topic string) error {
	table := getTopicTableName(topic)
	return Client.ExecDDL(fmt.Sprintf(createTopicDataTableSQL, table, table, table))
}

const insertTopicData = `
insert into %s (topic,offset,body) values (?, ? , ?)
`

func InsertTopicData(topic string, offset int64, body []byte) error {
	_, err := Client.Insert(fmt.Sprintf(insertTopicData, getTopicTableName(topic)), topic, offset, body)
	return err
}

func InsertTopicDataTransaction(topic string, offset int64, body []byte) error {
	err := Client.Transaction(func(tx *sql.Tx) error {
		stm, err := tx.Prepare(fmt.Sprintf(insertTopicData, getTopicTableName(topic)))
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
		stm, err = tx.Prepare(updateTopic)
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

const insertTopic = `
insert into topic_detail (topic) values (?)
`

func InsertTopic(topic string) error {
	_, err := Client.Insert(insertTopic, topic)
	return err
}

const updateTopic = `
update topic_detail set offset = ?,last_modified_time = datetime('now') where topic = ?
`

func UpdateTopic(topic string, offset int64) error {
	_, err := Client.Update(updateTopic, offset, topic)
	return err
}

const findByTopic = `
select * from topic_detail where topic = ?
`

func FindByTopic(topic string) (*model.TopicDetail, error) {
	var td *model.TopicDetail
	err := Client.FindFirst(findByTopic, &td, topic)
	return td, err
}
