package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hi")
}
