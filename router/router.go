package router

import (
	adminRouter "ginstart/app/admin/router"
	"ginstart/global"
	"ginstart/middleware"
	"io"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	runMode := global.Conf.System.RunMode
	gin.SetMode(runMode)

	routers := gin.Default()

	switch runMode {
	case "release":
		// 禁止日志颜色
		gin.DisableConsoleColor()
		// 禁止控制台输出
		gin.DefaultWriter = io.Discard
	case "debug":
		// 火焰图
		pprof.Register(routers)
	default:
		// 禁止颜色
		gin.DisableConsoleColor()
		// 禁止Gin的控制台输出
		gin.DefaultWriter = io.Discard
	}

	//允许修复当前请求路径，如/FOO和/..//Foo会被修复为/foo，并进行重定向，默认为 false。
	routers.RedirectFixedPath = true

	// //异常处理中间件
	// routers.Use(middleware.Recover)

	//跨域中间件
	routers.Use(middleware.Cors())
	//实例化路由
	systemRouter := adminRouter.SystemRouter{}

	//授权路由
	privateGroup := routers.Group("/")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.UberRateLimit())
	{
		systemRouter.InitPrivateRouter(privateGroup) //系统路由
	}

	//公开路由
	publicGroup := routers.Group("")
	publicGroup.Use(middleware.UberRateLimit())
	{
		systemRouter.InitPublicRouter(publicGroup) //系统路由
	}

	return routers
}
