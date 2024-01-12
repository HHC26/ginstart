package db

import (
	"ginstart/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormPgSQLTool struct{}

var GormPgSQL = new(GormPgSQLTool)

// InitDB 初始化数据库
func (g *GormPgSQLTool) InitDB() *gorm.DB {
	sql := global.Conf.Mysql
	dsn := sql.Username + ":" + sql.Password + "@tcp(" + sql.Path + ")/" + sql.DbName + "?charset=" + sql.Config

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		global.Log.Error("连接Pgsql数据库出现错误：", err)
		panic("连接Pgsql数据库出现错误：" + err.Error())
	}

	// 开启sqlite日志
	if sql.LogMode {
		db.Debug()
	}

	return db
}

// GetTableInfoList 获取表信息
func (g *GormPgSQLTool) GetTableInfoList() (list []DBTableInfo, err error) {
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
func (g *GormPgSQLTool) GetColumnInfoList(tableName string) (list []DBColumnInfo, err error) {
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
