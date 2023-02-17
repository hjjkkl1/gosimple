package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/valyevo/gosimple/internal/routes"
)

func initConfig() {
	// 配置路径
	baseDir := "configs"
	path := baseDir + "/" + cfgFilename + ".yaml"

	// 读取匹配的环境变量
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "[config] config not found >>", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stdout, "[config] using config file >>", viper.ConfigFileUsed())
	}
}

func initGin(routef routes.RouteFunc) error {
	// 配置服务器模式
	gin.SetMode(viper.GetString("runmode"))
	// 实例化 Gin 引擎
	g := gin.New()
	// 使用默认中间件
	g.Use(gin.Logger(), gin.Recovery())
	// 加载其他中间件
	// g.Use()
	// 加载路由
	appgroup := g.Group(routeGroupName)
	routef(appgroup)
	// 创建 HTTP 服务实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	// 打印 Log
	fmt.Println("[server] server start")
	// 运行
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()
	// 创建退出等待信号
	quit := make(chan os.Signal, 1)
	// 等待给定的信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 等待退出信号
	<-quit

	// 创建 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 关闭服务
	if err := httpsrv.Shutdown(ctx); err != nil {
		// TODO: 处理错误日志
		return err
	}

	return nil
}
