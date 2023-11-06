package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/pkg/server/controller"
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
				return
			},
		},
		{
			Name:    "favicon.ico",
			Method:  "GET",
			Pattern: "/favicon.ico",
			HandlerFunc: func(ctx *gin.Context) {
				return
			},
		},
	}

	routes = []Route{
		{
			Name:        "routes",
			Method:      "GET",
			Pattern:     "hello",
			HandlerFunc: c.Hello,
			Middleware:  gin.HandlersChain{},
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
