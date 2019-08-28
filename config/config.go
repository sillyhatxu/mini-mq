package config

var Conf Config

type Config struct {
	Http Http
	Log  LogConf
}

type Http struct {
	Listen string
}

//log.dir
type LogConf struct {
	OpenLogstash    bool   `toml:"open_logstash"`
	OpenLogfile     bool   `toml:"open_logfile"`
	FilePath        string `toml:"file_path"`
	Project         string `toml:"project"`
	Module          string `toml:"module"`
	LogstashAddress string `toml:"logstash_address"`
}

type MQConf struct {
	AutoCreateTopicEnable bool  //是否允许自动创建topic
	DeleteTopicEnable     bool  //是否允许删除topic
	PartitionCount        int64 //数据库拆分记录数
	MaxGoroutine          int   //最大goroutin数量
}
