package logger

import "go.uber.org/zap"

func NewZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"./logs/dev09.log"}

	return zap.Must(cfg.Build())
}
