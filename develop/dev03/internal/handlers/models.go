package handlers

import (
	"dev03/internal/domain"
	"dev03/pkg/logger"

	"go.uber.org/zap"
)

type Handler struct {
	Log    *zap.Logger
	Domain *domain.Domain
}

func New() *Handler {
	return &Handler{
		Log:    logger.NewZap(),
		Domain: domain.New(),
	}
}
