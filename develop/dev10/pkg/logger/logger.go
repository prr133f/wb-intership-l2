package logger

import "go.uber.org/zap"

func NewZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	// cfg.Encoding = "json"
	// cfg.OutputPaths = []string{"./logs/dev10.log"}

	return zap.Must(cfg.Build())
}
