package handler

import (
	"main/package/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	handler *http.Handler
}

func (h *Handler) ServeRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/", service.HelloFunc)
	}
	return router
}
