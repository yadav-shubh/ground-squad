package utils

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

var Log = getLogger()
var doOnce sync.Once
var logger *zap.Logger

func getLogger() *zap.Logger {
	doOnce.Do(func() {
		l, err := zap.NewProduction()
		if err != nil {
			log.Panic("error getting logger:", zap.Error(err))
		}

		l.Info("Initialized logger..")
		logger = l
	})
	return logger
}
