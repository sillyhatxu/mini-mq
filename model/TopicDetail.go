package model

import "time"

type TopicDetail struct {
	Topic            string `json:"topic" mapstructure:"topic"`
	Offset           int64  `json:"offset" mapstructure:"offset"`
	CreatedTime      time.Time
	LastModifiedTime time.Time
	//RecordDelayTime time.Duration
}

//file.delete.delay.ms 文件保存时间
