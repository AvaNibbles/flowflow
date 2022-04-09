package internal

import "go.uber.org/zap"

func GetLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
