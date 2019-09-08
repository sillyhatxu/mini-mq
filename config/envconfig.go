package config

var Conf config

type config struct {
	HttpPort int `toml:"http_port"`
	GRPCPort int `toml:"grpc_port"`
	Project  string
	Module   string
	Log      logConf  `toml:"log_conf"`
	DB       database `toml:"database"`
}

type database struct {
	DataSourceName string `toml:"data_source_name"`
	DDLPath        string `toml:"ddl_path"`
}

type logConf struct {
	OpenLogstash    bool   `toml:"open_logstash"`
	OpenLogfile     bool   `toml:"open_logfile"`
	FilePath        string `toml:"file_path"`
	LogstashAddress string `toml:"logstash_address"`
}

//type MQConf struct {
//	AutoCreateTopicEnable bool  //是否允许自动创建topic
//	DeleteTopicEnable     bool  //是否允许删除topic
//	PartitionCount        int64 //数据库拆分记录数
//	MaxGoroutine          int   //最大goroutin数量
//}
