package router

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"app/pkg/conf"
	"app/pkg/logger"
	"app/pkg/server/router/middleware"
	"app/pkg/util"
)

var routeGroup []*GroupRoute

func Server(cfg *conf.ConfigYaml) *http.Server {
	//gin.DefaultWriter = logger.L
	routeConfig()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port),
		Handler: setupRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.L.Error(err.Error())
			os.Exit(1)
		}
	}()
	logger.L.Info(fmt.Sprintf("start on server:%s", srv.Addr))
	return srv
}

func GracefulExit(srv *http.Server, ch chan struct{}) {
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.L.Error(err.Error())
	}
	logger.L.Info("Shutdown server ...")
}

func routeConfig() {
	var commonRoute []Route
	commonRoute = append(commonRoute, routes...)
	routeGroup = []*GroupRoute{
		{
			Prefix: "/api/v1",
			GroupMiddleware: gin.HandlersChain{
				middleware.RecoveryWithWriter(),
			},
			SubRoutes: commonRoute,
		},
		{
			Prefix:    "",
			SubRoutes: baseRoutes,
		},
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors(),
		middleware.LogRequest(),
		middleware.GinLogger(logger.L),
		middleware.GinRecovery(logger.L, true))
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// sso
	gob.Register(map[string]any{}) // important!
	store := cookie.NewStore([]byte(viper.GetString("gateway.secretKey")))
	r.Use(sessions.Sessions("session", store))

	routeGroupsMap := make(map[string]*gin.RouterGroup)
	for _, gRoute := range routeGroup {
		if _, ok := routeGroupsMap[gRoute.Prefix]; !ok {
			routeGroupsMap[gRoute.Prefix] = r.Group(gRoute.Prefix)
		}
		for _, gMiddleware := range gRoute.GroupMiddleware {
			routeGroupsMap[gRoute.Prefix].Use(gMiddleware)
		}
		for _, subRoute := range gRoute.SubRoutes {
			length := len(subRoute.Middleware) + 2
			routes := make([]any, length)
			routes[0] = subRoute.Pattern
			for i, v := range subRoute.Middleware {
				routes[i+1] = v
			}
			routes[length-1] = subRoute.HandlerFunc

			util.CallReflect(
				routeGroupsMap[gRoute.Prefix],
				subRoute.Method,
				routes...)
		}
	}
	r.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	return r
}
