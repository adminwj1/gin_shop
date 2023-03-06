package v1

import (
	"gin_demo/internal/controller/backend"

	"github.com/gin-gonic/gin"
)

func BackendRouter(router *gin.RouterGroup) {
	router.GET("/rotation", backend.Rotation)
}
