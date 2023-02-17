package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valyevo/gosimple/internal/pkg/core"
	"github.com/valyevo/gosimple/internal/pkg/errcode"
)

type RouteFunc func(app *gin.RouterGroup)

func ApiRoutes(app *gin.RouterGroup) {
	// 注册应用中间件
	app.Use()
	// 注册路由
	app.GET("", func(ctx *gin.Context) {
		data := []interface{}{
			"well",
			"hello",
			"valye",
		}

		core.WriteResponse(ctx, errcode.OK, data)
	})
}
