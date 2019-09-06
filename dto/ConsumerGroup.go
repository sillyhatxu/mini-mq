package dto

type ConsumerGroup struct {
	TopicName  string `json:"topic_name"`
	TopicGroup string `json:"topic_group"`
	Offset     int64  `json:"offset"`
	Body       []byte `json:"body"`
}
