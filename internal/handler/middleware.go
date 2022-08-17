package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")

	if ctx.Request.Method != "OPTIONS" {
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusOK)
	}
}
