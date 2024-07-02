package domain

import (
	"go.uber.org/zap"
)

type Domain struct {
	Log *zap.Logger
}

func NewDomain(logger *zap.Logger) *Domain {
	return &Domain{
		Log: logger,
	}
}

type Flags struct {
	B string
	C string
	D string
	F string
	S bool
}

type parsedRange struct {
	first  int
	last   int
	isSolo bool
}
