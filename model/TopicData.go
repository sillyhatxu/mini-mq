package model

type TopicData struct {
	TopicName string `json:"topic_name" mapstructure:"topic_name"`
	Offset    int64  `json:"offset" mapstructure:"offset"`
	Body      []byte `json:"body" mapstructure:"body"`
}
