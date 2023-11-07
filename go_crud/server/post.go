package server

import (
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
)

func AddPOST(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/add", func(c *gin.Context) {
		var listRes mysql_db.CrudList
		err := c.ShouldBindJSON(&listRes) //数据校验
		if err != nil {                   //数据错
			c.JSON(200, gin.H{
				"msg":  "添加失败",
				"data": err,
				"code": "400",
			})
		} else {
			// 数据库操作，insert
			db.Create(&listRes)
			c.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": listRes,
				"code": "200",
			})
		}
	})
	return
}
