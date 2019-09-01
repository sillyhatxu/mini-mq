package model

type TopicGroup struct {
	TopicName  string `json:"topic_name" mapstructure:"topic_name"`
	TopicGroup string `json:"topic_group" mapstructure:"topic_group"`
	Offset     int64  `json:"offset" mapstructure:"offset"`
}
