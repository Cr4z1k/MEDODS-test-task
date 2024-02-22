package handler

import (
	"github.com/Cr4z1k/MEDODS-test-task/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	token := r.Group("/token")
	{
		token.Handle("POST", "/new/:guid", h.GetTokens)
		token.Handle("POST", "/refresh", h.RefreshAccess)
	}

	return r
}
