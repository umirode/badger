package logger

import "github.com/sirupsen/logrus"

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

	logrus.Error(message+" ", err.Error())
}

func (l *LogrusLogger) LogInfo(message string) {
	logrus.Info(message)
}
