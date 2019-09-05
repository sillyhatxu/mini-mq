package consumer

type Group struct {
	TopicName      string
	TopicGroup     string
	LatestOffset   int64
	TopicDataArray []TopicData
}

type TopicData struct {
	TopicName  string
	TopicGroup string
	Offset     int64
	Body       []byte
}
