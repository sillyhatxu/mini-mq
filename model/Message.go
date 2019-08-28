package model

//m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value)
type Message struct {
	Topic  string
	Offset int64
	Body   []byte
}
