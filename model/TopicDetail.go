package model

import "time"

type TopicDetail struct {
	TopicName        string `json:"topic_name" mapstructure:"topic_name"`
	Offset           int64  `json:"offset" mapstructure:"offset"`
	CreatedTime      time.Time
	LastModifiedTime time.Time
	//RecordDelayTime time.Duration
}

//file.delete.delay.ms 文件保存时间
