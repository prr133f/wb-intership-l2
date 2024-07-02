package domain

import "go.uber.org/zap"

type Domain struct {
	Log *zap.Logger
}

func NewDomain(logger *zap.Logger) *Domain {
	return &Domain{
		Log: logger,
	}
}
