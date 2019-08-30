package service

import (
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	dbclient.InitialDBClient("/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	err := CreateTopic("test_topic")
	assert.Nil(t, err)
}

func TestCreateTopic2(t *testing.T) {
	dbclient.InitialDBClient("/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	err := CreateTopic("test_topic2")
	assert.Nil(t, err)
}
