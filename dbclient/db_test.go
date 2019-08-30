package dbclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	InitialDBClient("/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	err := Client.Initial()
	if err != nil {
		panic(err)
	}
}

func TestUpdateTime(t *testing.T) {
	_, err := Client.Update("update schema_version set created_time = datetime('now') where id = 4")
	assert.Nil(t, err)
}

func TestUpdateTimeLocal(t *testing.T) {
	_, err := Client.Update("update schema_version set created_time = datetime(CURRENT_TIMESTAMP,'localtime') where id = 5")
	assert.Nil(t, err)
}

func TestFindByTopic(t *testing.T) {
	topicDetail, err := FindByTopic("test")
	assert.Nil(t, err)
	assert.Nil(t, topicDetail)
}
