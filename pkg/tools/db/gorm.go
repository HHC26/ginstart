package db

import (
	"fmt"
	"ginstart/app/admin/models"
	"ginstart/global"

	"gorm.io/gorm"
)

type GormTool struct{}

var Gorm = new(GormTool)

type Database interface {
	InitDB() *gorm.DB
	GetTableInfoList() (list []DBTableInfo, err error)
	GetColumnInfoList(tableName string) (list []DBColumnInfo, err error)
}

func (gt *GormTool) Database() Database {
	switch global.Conf.System.DbType {
	case "mysql":
		return GormMySQL
	case "pgsql":
		return GormPgSQL
	case "sqlite":
		return GormSqlite
	default:
		return GormMySQL
	}
}

func (gt *GormTool) InitTables(db *gorm.DB) {
	// 数据迁移
	err := db.AutoMigrate(
		&models.SysUser{},
		&models.SysRole{},
	)

	if err != nil {
		fmt.Println("Data migration failed, err:", err)
		global.Log.Error("Data migration failed, err:", err)
	}
}

type DBTableInfo struct {
	TableName        string `json:"tableName"`        //表名称
	TableDescription string `json:"tableDescription"` //表描述
	SchemaName       string `json:"schemaName"`       //pgsql架构名称
}

type DBColumnInfo struct {
	TableName         string `json:"tableName"`         //表名称
	ColumnName        string `json:"columnName"`        //字段名称
	ColumnDescription string `json:"columnDescription"` //字段描述
	ColumnType        string `json:"columnType"`        //字段类型
	IsPk              bool   `json:"isPk"`              //是否主键
}
