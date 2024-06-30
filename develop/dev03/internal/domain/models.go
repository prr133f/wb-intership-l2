package domain

import (
	"dev03/pkg/logger"

	"go.uber.org/zap"
)

type Domain struct {
	Log *zap.Logger
}

func New() *Domain {
	return &Domain{
		Log: logger.NewZap(),
	}
}

type Flags struct {
	K int
	N bool
	R bool
	U bool
}
