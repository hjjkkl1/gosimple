package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteFunc func(app *gin.RouterGroup)

func ApiRoutes(app *gin.RouterGroup) {
	// 注册应用中间件
	app.Use()
	// 注册路由
	app.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello",
			"code":    0,
		})
	})
}
