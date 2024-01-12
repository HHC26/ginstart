package router

import (
	"ginstart/app/admin/apis"

	"github.com/gin-gonic/gin"
)

// SystemRouter 系统路由
// 建议：
// 当接口数量超出一定范围，可适当分类或将复杂模块独立文件整理
// 路由规范：业务/模块/操作，比如：mall/user/create
type SystemRouter struct{}

// var (
// 	authApi           = &api.SysAuthApi{}
// 	configApi         = &api.SysConfigApi{}
// 	menuApi           = &api.SysMenuApi{}
// 	userApi           = &api.SysUserApi{}
// 	roleApi           = &api.SysRoleApi{}
// 	genTableApi       = &api.SysGenTableApi{}
// 	genTableColumnApi = &api.SysGenTableColumnApi{}
// 	monitorApi        = &api.SysMonitorApi{}
// 	logApi            = &api.SysLogApi{}
// )

// InitPublicRouter 初始化公开路由
func (s *SystemRouter) InitPublicRouter(routerGroup *gin.RouterGroup) {
	authRouter := routerGroup.Group("/system")
	{
		authRouter.GET("heatch_check", apis.HeathCheck)
	}

}

// InitPrivateRouter 初始化私有路由
func (s *SystemRouter) InitPrivateRouter(routerGroup *gin.RouterGroup) {

}
