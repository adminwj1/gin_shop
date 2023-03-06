package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 测试接口
// @Success 200 {string} string json{"code","message"}
// @Router  /backend/rotation [get]
// @version 1.0
func Rotation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK！",
	})
}
