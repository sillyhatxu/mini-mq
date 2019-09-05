package producer

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_Produce(t *testing.T) {
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
	userinfo := &UserInfo{Id: fmt.Sprintf("ID_%v", 1), MobileNumber: "555555", Name: fmt.Sprintf("test-%v", 1), Paid: true, FirstActionDeviceId: "deviceid", TestNumber: 10, TestNumber64: 64, TestDate: time.Now()}
	userJSON, err := json.Marshal(userinfo)
	assert.Nil(t, err)
	Client := &Client{Address: Address, Timeout: 60 * time.Second}
	err = Client.Produce(TopicName, userJSON)
	assert.Nil(t, err)
}
