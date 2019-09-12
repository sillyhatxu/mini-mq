package topic

import (
	"fmt"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

//const path = "/Users/shikuanxu/go/src/github.com/sillyhatxu/mini-mq"

const path = "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq"

func init() {
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	cache.Initial()
}

func TestCreateTopic(t *testing.T) {
	err := CreateTopic("test_topic")
	assert.Nil(t, err)

}

func TestCreateTopic2(t *testing.T) {
	err := CreateTopic("test_topic2")
	assert.Nil(t, err)
}
