package config

import (
	"github.com/sillyhatxu/logrus-client"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
)

func InitialLogConfig() {
	fields := logrus.Fields{
		"project":  Conf.Project,
		"module":   Conf.Module,
		"@version": "1",
		"type":     "project-log",
	}
	var fileConfig = filehook.NewFileConfig(Conf.Log.FilePath)
	var logstashConfig = logstashhook.NewLogstashConfig(Conf.EnvConfig.LogstashURL, logstashhook.Fields(fields))
	var config = logrusconf.NewLogrusConfig(
		logrusconf.Fields(fields),
		logrusconf.FileConfig(fileConfig),
		logrusconf.LogstashConfig(logstashConfig),
	)
	config.Initial()
}
