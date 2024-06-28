package domain

import "go.uber.org/zap"

type Domain struct {
	Log *zap.Logger
}

func New() *Domain {
	return &Domain{
		Log: zap.NewExample(),
	}
}

type Flags struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}
