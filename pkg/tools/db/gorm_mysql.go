package db

import (
	"ginstart/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormMySQLTool struct{}

var GormMySQL = new(GormMySQLTool)

// InitDB 初始化数据库
func (g *GormMySQLTool) InitDB() *gorm.DB {
	sql := global.Conf.Mysql
	dsn := sql.Username + ":" + sql.Password + "@tcp(" + sql.Path + ")/" + sql.DbName + "?charset=" + sql.Config

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "la_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		global.Log.Error("连接Mysql数据库出现错误：", err)
		panic("连接Mysql数据库出现错误：" + err.Error())
	}

	// 开启sqlite日志
	if sql.LogMode {
		db.Debug()
	}

	// 新增内容
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("initMysql db.DB err:", err)
		panic("initMysql db.DB err:" + err.Error())
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(sql.MaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(sql.MaxOpenConns)

	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(2) * time.Hour)

	return db
}

// GetTableInfoList 获取表信息
func (g *GormMySQLTool) GetTableInfoList() (list []DBTableInfo, err error) {
	tableSql := `
		SELECT 
			table_name AS TableName,
			table_comment AS TableDescription
		FROM 
			information_schema.TABLES 
		WHERE 
			table_schema = (SELECT DATABASE())`
	err = global.Db.Raw(tableSql).Find(&list).Error
	return list, err
}

// GetColumnInfoList 获取表信息
func (g *GormMySQLTool) GetColumnInfoList(tableName string) (list []DBColumnInfo, err error) {
	columnSql := `		
	SELECT
		table_name AS TableName,
		column_name AS ColumnName,
		column_comment AS ColumnDescription,
		data_type AS ColumnType,
		( CASE WHEN column_key = 'PRI' THEN true ELSE false END ) AS IsPk 
	FROM
		information_schema.COLUMNS 
	WHERE
		table_schema = (SELECT DATABASE()) AND table_name = ? ORDER BY ORDINAL_POSITION`
	err = global.Db.Raw(columnSql, tableName).Find(&list).Error
	return list, err
}
