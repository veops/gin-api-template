package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/pkg/logger"
	"app/pkg/server/controller"
	"app/pkg/server/router/middleware"
)

var (
	c = controller.NewController()

	baseRoutes = []Route{
		{
			Name:    "a health check, just for monitoring",
			Method:  "GET",
			Pattern: "/-/health",
			HandlerFunc: func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "OK")
			},
		},
		{
			Name:    "favicon.ico",
			Method:  "GET",
			Pattern: "/favicon.ico",
			HandlerFunc: func(ctx *gin.Context) {
			},
		},
		{
			Name:    "change the log level",
			Method:  "PUT",
			Pattern: "/-/log/level",
			HandlerFunc: func(ctx *gin.Context) {
				logger.AtomicLevel.ServeHTTP(ctx.Writer, ctx.Request)
			},
		},
	}

	routes = []Route{
		{
			Name:        "routes",
			Method:      "GET",
			Pattern:     "hello",
			HandlerFunc: c.Hello,
			Middleware:  gin.HandlersChain{middleware.Auth()},
		},
	}
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(ctx *gin.Context)
	Middleware  []gin.HandlerFunc
}

type GroupRoute struct {
	Prefix          string
	GroupMiddleware gin.HandlersChain
	SubRoutes       []Route
}
