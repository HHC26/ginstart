package start

import (
	"context"
	"fmt"
	"ginstart/cmd/version"
	"ginstart/docs"
	"ginstart/global"
	"ginstart/pkg/tools/db"
	"ginstart/pkg/tools/log"
	"ginstart/router"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "start -c config/config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

// 初始化swagger文档
func initSwagger(routers *gin.Engine) {
	docs.SwaggerInfo.Title = global.Conf.System.AppName
	docs.SwaggerInfo.Description = "Swagger Admin API"
	docs.SwaggerInfo.Version = version.Version
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}
	//初始化配置
	global.Conf = global.InitConfig(configYml)

	// 初始化日志
	global.Log = log.InitLogger()
	// 初始化数据库连接
	global.Db = db.Gorm.Database().InitDB()
	// 迁移数据表
	db.Gorm.InitTables(global.Db)
	// 初始化路由
	routers := router.InitRouters()
	initSwagger(routers)

	system := global.Conf.System
	//自定义 HTTP 服务器
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(system.Port),
		Handler:        routers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf(`	欢迎使用 %s %s
	数据库类型[%s] ,日志类型[%s]
	访问地址 http://127.0.0.1%s
	`, system.AppName, version.Version, system.DbType, global.Conf.Zap.Level, server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error("Server listen err: %s\n", err)
		}
	}()

	shutDown(server)
}

func shutDown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Error("Server Shutdown: ", err)
	}
	global.Log.Info("Server Exit")
}
