package db

import (
	"fmt"
	"ginstart/global"
	"ginstart/pkg/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormSqliteTool struct{}

var GormSqlite = new(GormSqliteTool)

// InitDB 初始化数据库
func (g *GormSqliteTool) InitDB() *gorm.DB {
	sql := global.Conf.Sqlite
	dbFile := fmt.Sprintf("%s.db", sql.DbName)
	if utils.CheckFileNotExist(dbFile) {
		if err := utils.CreateFile(dbFile); err != nil {
			global.Log.Error("创建数据库文件失败：", err)
			panic("创建数据库文件失败：" + err.Error())
		}
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		// 表名命名策略
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		global.Log.Error("连接Sqlite数据库出现错误：", err)
		panic("连接Sqlite数据库出现错误：" + err.Error())
	}

	// 开启sqlite日志
	if sql.LogMode {
		db.Debug()
	}

	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")

	return db
}

// GetTableInfoList 获取表信息
func (g *GormSqliteTool) GetTableInfoList() (list []DBTableInfo, err error) {
	sql := `
		SELECT
			schemaname AS SchemaName,
			tablename AS TableName,
			CAST ( obj_description ( C.oid, 'pg_class' ) AS VARCHAR ) AS TableDescription
		FROM
			pg_tables AS tables
			INNER JOIN pg_class C ON tables.tablename = C.relname 
		WHERE
			schemaname NOT IN ( 'pg_catalog', 'information_schema' ) 
		ORDER BY schemaname,tablename`
	err = global.Db.Raw(sql).Find(&list).Error
	return list, err
}

// GetColumnInfoList 获取表信息
func (g *GormSqliteTool) GetColumnInfoList(tableName string) (list []DBColumnInfo, err error) {
	sql := `
		SELECT 
			table_name AS TableName,
			table_comment AS TableDescription
		FROM 
			information_schema.TABLES 
		WHERE 
			table_schema = (SELECT DATABASE())`
	err = global.Db.Raw(sql).Find(&list).Error
	return list, err
}
