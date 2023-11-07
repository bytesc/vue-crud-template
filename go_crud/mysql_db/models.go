package mysql_db

import "gorm.io/gorm"

// CrudList gorm结构体
type CrudList struct {
	gorm.Model // 引入模板结构体
	//ID        uint `gorm:"primarykey"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt DeletedAt `gorm:"index"`
	Name     string `gorm:"type: varchar(20); not null" binding:"required" json:"name" ` //binding:"required" 表示not bull
	Level    string `gorm:"type: varchar(20);" json:"level"`
	Email    string `gorm:"type: varchar(50);" json:"email"`
	Phone    string `gorm:"type: varchar(20);" json:"phone"`
	Birthday string `gorm:"type: varchar(20);" json:"birthday"`
	Address  string `gorm:"type: varchar(200);" json:"address"`
	// 变量名要大写，才public，可以被gorm获取解析
}
