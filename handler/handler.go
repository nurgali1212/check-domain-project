package handler

import (
	"domain/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	r := router.Group("/api")
	{
		r.POST("/check", h.checkDomain)

	}

	return router
}
