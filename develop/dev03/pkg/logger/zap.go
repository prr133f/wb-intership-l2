package logger

import (
	"go.uber.org/zap"
)

func NewZap() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())

	return logger
}
