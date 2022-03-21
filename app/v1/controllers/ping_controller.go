package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

type PingControllerInterface interface {
	Pong(c *gin.Context)
}

func NewPingController() PingControllerInterface {
	return &PingController{}
}

func (c *PingController) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
