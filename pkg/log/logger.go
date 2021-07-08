package log

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})

	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	*logrus.Logger
}

func New() Logger {
	return logrus.New()
}
