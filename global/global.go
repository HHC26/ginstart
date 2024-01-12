package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Conf *AppConfig         // 配置
	Log  *zap.SugaredLogger //日志
	Db   *gorm.DB           // 数据库
)
