package service

import (
	"encoding/json"
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

func TestProduce(t *testing.T) {
	dbclient.InitialDBClient("/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/basic.db", "/Users/cookie/go/gopath/src/github.com/sillyhatxu/mini-mq/db/migration")
	cache.Initial()
	userinfo := &UserInfo{Id: "ID_1001", MobileNumber: "555555", Name: "test", Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now()}
	userJSON, err := json.Marshal(userinfo)
	assert.Nil(t, err)
	err = Produce("test_topic", userJSON)
	assert.Nil(t, err)
}
