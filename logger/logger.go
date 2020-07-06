package logger

type ILogger interface {
	LogError(message string, err error)
	LogInfo(message string)
}
