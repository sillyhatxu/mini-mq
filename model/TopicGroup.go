package model

type TopicGroup struct {
	Topic         string
	Group         string
	Offset        int64
	ConsumerCount int
}
