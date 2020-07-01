package main

import "github.com/sirupsen/logrus"

type ILogger interface {
	LogError(message string, err error)
}

type LogrusLogger struct {
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (l *LogrusLogger) LogError(message string, err error) {
	if err == nil {
		logrus.Error(message)
		return
	}

	logrus.Error(message, err.Error())
}
