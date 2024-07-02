package handlers

import (
	"dev06/internal/domain"
	"dev06/pkg/logger"

	"go.uber.org/zap"
)

type Handler struct {
	Log    *zap.Logger
	Domain *domain.Domain
}

func NewHandler() *Handler {
	log := logger.NewZap()
	return &Handler{
		Log:    log,
		Domain: domain.NewDomain(log),
	}
}
