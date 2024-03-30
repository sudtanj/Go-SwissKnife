package log

import "go.uber.org/zap"

type LoggerMap map[string]interface{}

func ToMapInterface(log zap.Field) LoggerMap {
	return LoggerMap{
		log.Key: log.Interface,
	}
}
