package producer

import (
	"encoding/json"
	"fmt"
	"github.com/sillyhatxu/mini-mq/dbclient"
	"github.com/sillyhatxu/mini-mq/utils/cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type UserInfo struct {
	Id                  string    `json:"id" mapstructure:"id"`
	MobileNumber        string    `json:"mobile_number" mapstructure:"mobile_number"`
	Name                string    `json:"Name" mapstructure:"Name"`
	Paid                bool      `json:"Paid" mapstructure:"Paid"`
	FirstActionDeviceId string    `json:"first_action_device_id" mapstructure:"first_action_device_id"`
	TestNumber          int       `json:"test_number" mapstructure:"test_number"`
	TestNumber64        int64     `json:"test_number_64" mapstructure:"test_number_64"`
	TestDate            time.Time `json:"test_date" mapstructure:"test_date"`
}

//const path = "/Users/shikuanxu/go/src/github.com/sillyhatxu/mini-mq"

const path = "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq"

func TestProduce(t *testing.T) {
	//dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), "")
	cache.Initial()
	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now()}
	userJSON, err := json.Marshal(userinfo)
	assert.Nil(t, err)
	err = Produce("test_topic", userJSON)
	assert.Nil(t, err)
}

//10W data  50.537781494s
func TestProduceBatch(t *testing.T) {
	//dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), fmt.Sprintf("%s/db/migration", path))
	dbclient.InitialDBClient(fmt.Sprintf("%s/basic.db", path), "")
	cache.Initial()
	start := time.Now()
	for i := 1; i < 100000; i++ {
		userinfo := &UserInfo{Id: fmt.Sprintf("ID_%v", i), MobileNumber: "555555", Name: fmt.Sprintf("test-%v", i), Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now()}
		userJSON, err := json.Marshal(userinfo)
		assert.Nil(t, err)
		err = Produce("test_topic", userJSON)
		assert.Nil(t, err)
	}
	fmt.Println(time.Since(start))
}
