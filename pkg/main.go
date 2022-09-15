package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nginx_mate_go/pkg/api"
	"nginx_mate_go/pkg/nginx"
	"nginx_mate_go/pkg/util"
)

func main() {
	nginx.Start()
	router := gin.Default()
	router.POST("/v1/nginx/route/static/tcp", api.AddStaticRoute)
	router.DELETE("/v1/nginx/route/static/tcp/:port", api.DelStaticRoute)
	util.Logger().Info("nginx mate started")
	err := router.Run(":31014")
	if err != nil {
		util.Logger().Error("open http server failed", zap.Error(err))
		return
	}
}
