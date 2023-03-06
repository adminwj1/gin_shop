package router

import (
	v1 "gin_demo/api/v1"
	_ "gin_demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter() *gin.Engine {
	e := gin.Default()
	apiGroup := e.Group("/frontend")
	backendGroup := e.Group("/backend")

	v1.BackendRouter(backendGroup)
	v1.FrontendRouter(apiGroup)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}

func RunServer() {
	r := SetUpRouter()
	r.Run(":8080")
}
