package handlers

import (
	"dev10/internal/domain"
	"dev10/pkg/logger"

	"go.uber.org/zap"
)

type Handler struct {
	Domain *domain.Domain
	Log    *zap.Logger
}

func NewHandler() *Handler {
	log := logger.NewZap()
	return &Handler{
		Log:    log,
		Domain: domain.NewDomain(log),
	}
}
