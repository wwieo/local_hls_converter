package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
	_ "main/docs"
	"main/service/internal/config"
	"net/http"
)

type servicePack struct {
	dig.In

	ServicePort config.ServicePort
	Handler     *gin.Engine
}

func NewServer(pack servicePack) *http.Server {
	return &http.Server{
		Addr:    pack.ServicePort.String(),
		Handler: pack.Handler,
	}
}

func NewRouterRoot(pack servicePack) *gin.RouterGroup {
	return pack.Handler.Group("hls_provider")
}

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.ContextWithFallback = true
	return router
}

// NewBasic
// @Tags		Surprise Checker
// @version	1.0
// @produce	text/plain
// @Success	200
// @Router		/Ping [GET]
func NewBasic(pack basicPack) {
	pack.Root.GET("Ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})
	pack.Root.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

type basicPack struct {
	dig.In

	Root *gin.RouterGroup
}
