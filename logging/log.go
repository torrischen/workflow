package logging

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewProduction()
	logger = l.Sugar()
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args)
}
