package handlers

import (
	"dev05/internal/domain"

	"go.uber.org/zap"
)

type Handlers struct {
	Log    *zap.Logger
	Domain *domain.Domain
}

func New() *Handlers {
	return &Handlers{
		Log:    zap.NewExample(),
		Domain: domain.New(),
	}
}
