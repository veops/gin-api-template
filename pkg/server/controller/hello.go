package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"app/pkg/server/model"
)

type HelloResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    model.Hello `json:"data"`
}

// Hello godoc
// @Tags Hello
// @Summary get Hello information
// @Accept text/plain
// @Produce json
// @Success 200 {object} HelloResponse
// @Failure default {object} HTTPError
// @Router /hello [get]
func (c *Controller) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HelloResponse{Code: http.StatusOK, Message: "hello world", Data: model.Hello{Time: time.Now()}})
	//ctx.JSON(200, gin.H{"code": 0, "message": "hello world", "data": model.Hello{Time: time.Now()}})
}
