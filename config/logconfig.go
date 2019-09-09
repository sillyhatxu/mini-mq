package config

import (
	"github.com/sillyhatxu/logrus-client"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
	"time"
)

func InitialLogConfig() {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	textFormatter := &logrus.TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		TimestampFormat:  string("2006-01-02 15:04:05"),
		QuoteEmptyFields: true,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "timestamp",
		},
	}
	conf := &logrusconf.Conf{
		Level:        logrus.InfoLevel,
		ReportCaller: true,
		Fields: logrus.Fields{
			"project": Conf.Project,
			"module":  Conf.Module,
		},
		LogFormatter: textFormatter,
	}
	if Conf.Log.OpenLogfile {
		conf.FileConf = &filehook.FileConf{
			LogFormatter:     textFormatter,
			FilePath:         Conf.Log.FilePath,
			WithMaxAge:       time.Duration(876000) * time.Hour,
			WithRotationTime: time.Duration(24) * time.Hour,
		}
	}
	if Conf.Log.OpenLogstash {
		conf.LogstashConf = &logstashhook.LogstashConf{
			LogFormatter: jsonFormatter,
			Address:      Conf.EnvConfig.LogstashURL,
		}
	}
	conf.Initial()
}
