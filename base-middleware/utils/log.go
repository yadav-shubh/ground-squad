package utils

import "go.uber.org/zap"

var Log = getLogger()

func initLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func getLogger() *zap.Logger {
	logger, err := initLogger()
	if err != nil {
		panic(err)
	}
	return logger
}
