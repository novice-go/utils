package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLog() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	Logger = l
}
