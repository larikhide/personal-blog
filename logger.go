package main

import (
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
)

type ConfigLogger struct {
	Level  string `yaml:"level"`
	Syslog bool   `yaml:"syslog"`
	Output string `yaml:"output"`
}

func ConfigureLogger(config *ConfigLogger) (*logrus.Logger, error) {
	lg := logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}
	lg.SetLevel(level)

	if config.Syslog {
		syslogFile, err := syslog.New(syslog.LOG_DEBUG, "personal-blog")
		if err != nil {
			return nil, err
		}
		lg.SetOutput(syslogFile)
	} else {
		if config.Output != "" {
			file, err := os.Create(config.Output)
			if err != nil {
				return nil, err
			}
			lg.SetOutput(file)
		}
	}
	return lg, nil
}
