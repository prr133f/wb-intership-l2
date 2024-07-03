package logger

import "go.uber.org/zap"

func NewAppZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	// cfg.Encoding = "json"
	// cfg.OutputPaths = []string{"./logs/dev11.log"}

	return zap.Must(cfg.Build())
}

func NewReqZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	// cfg.Encoding = "json"
	// cfg.OutputPaths = []string{"./logs/dev11.log"}

	return zap.Must(cfg.Build())
}
