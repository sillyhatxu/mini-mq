package dto

//m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value)
//type Message struct {
//	Topic  string
//	Offset int64
//	Body   []byte
//}

type TopicData struct {
	TopicName  string `json:"topic_name"`
	TopicGroup string `json:"topic_group"`
	Offset     int64  `json:"offset"`
	Body       []byte `json:"body"`
}

type TopicDataDTO struct {
	TopicName string `json:"topic_name"`
	Offset    int64  `json:"offset"`
	Body      string `json:"body"`
}
