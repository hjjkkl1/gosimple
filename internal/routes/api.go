package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valyevo/gosimple/internal/controller/user"
	"github.com/valyevo/gosimple/internal/pkg/core"
	"github.com/valyevo/gosimple/internal/pkg/errcode"
	"github.com/valyevo/gosimple/internal/store"
)

type RouteFunc func(app *gin.RouterGroup)

// mws 注册中间件
var middlewares = []gin.HandlerFunc{
	cors.Default(),
}

// ApiRoutes 路由组
func ApiRoutes(app *gin.RouterGroup) {
	// 注册应用中间件
	app.Use(middlewares...)
	// 注册路由
	app.GET("", func(ctx *gin.Context) {
		data := []interface{}{
			"well",
			"hello",
			"valye",
		}

		core.WriteResponse(ctx, errcode.OK, data)
	})

	// User 注册用户路由
	userCtrl := user.New(store.D)
	userGrp := app.Group("/user")
	{
		// 创建用户
		userGrp.GET("", userCtrl.Create)
	}
}
