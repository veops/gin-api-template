package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"app/pkg/server/model"
)

func (c *Controller) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"code": 0, "message": "hello world", "data": model.Hello{Time: time.Now()}})
}
