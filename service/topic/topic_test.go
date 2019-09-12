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
const dataSourceName = "sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat_minimq?loc=Asia%2FSingapore&parseTime=true"

func init() {
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	cache.Initial()
	dbclient.InitialDBClient(dataSourceName, fmt.Sprintf("%s/db/migration", path))
}

func TestCreateTopic(t *testing.T) {
	err := CreateTopic("test_topic")
	assert.Nil(t, err)

}

func TestCreateTopic2(t *testing.T) {
	err := CreateTopic("test_topic2")
	assert.Nil(t, err)
}
