package dto

type ConsumingDTO struct {
	ConsumingKey string `json:"consuming_key"`
	Address      string `json:"address"`
	Status       string `json:"status"`
}
