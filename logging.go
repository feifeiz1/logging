package logging

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func Init() {
	loggers, _ := zap.NewDevelopment(zap.AddCallerSkip(1))
	logger = loggers.Sugar()
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args)
}
