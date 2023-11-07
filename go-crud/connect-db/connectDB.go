package connect_db

//package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// CrudList gorm结构体
type CrudList struct {
	gorm.Model // 引入模板结构体
	//ID        uint `gorm:"primarykey"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt DeletedAt `gorm:"index"`
	Name     string `gorm:"type: varchar(20); not null" binding:"required" json:"name" `
	Level    string `gorm:"type: varchar(20);" json:"level"`
	Email    string `gorm:"type: varchar(50);" json:"email"`
	Phone    string `gorm:"type: varchar(20);" json:"phone"`
	Birthday string `gorm:"type: varchar(20);" json:"birthday"`
	Address  string `gorm:"type: varchar(200);" json:"address"`
	// 变量名要大写，才public，可以被gorm获取解析
}

func ConnectToDatabase() (*gorm.DB, error) {
	//链接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名单数
		},
	})
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize:         256,                                                                                  // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                                                                                // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 数据库迁移
	err = db.AutoMigrate(&CrudList{})
}
