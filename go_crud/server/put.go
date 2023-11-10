package server

import (
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
	"strconv"
)

func UpdatePOST(r *gin.Engine, db *gorm.DB) {
	r.POST("api/user/update/:id/", func(c *gin.Context) {
		db = db.Session(&gorm.Session{NewDB: true}) //必须清空上次遗留的链式条件
		var data mysql_db.CrudList
		id := c.Param("id") //接收路径参数
		// c.Query()  //接收查询参数
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "id格式错误",
				"data": err,
				"code": "400",
			})
		} else {
			db.Select("id").Where("id = ?", idInt).Find(&data) //数据库查找
			if data.ID == 0 {
				c.JSON(200, gin.H{
					"msg":  "修改失败，数据不存在",
					"data": data,
					"code": "400",
				})
			} else {
				err := c.ShouldBindJSON(&data)
				if err != nil {
					c.JSON(200, gin.H{
						"msg":  "修改失败，数据格式不正确",
						"data": err.Error(),
						"code": "400",
					})
				} else {
					result := db.Where("id = ?", idInt).Updates(&data)
					if result.RowsAffected == 0 {
						c.JSON(200, gin.H{
							"msg":  "修改失败，零行变动",
							"data": data,
							"code": "400",
						})
					} else if result.Error != nil {
						c.JSON(200, gin.H{
							"msg":  "修改失败",
							"data": result.Error,
							"code": "400",
						})
					} else {
						c.JSON(200, gin.H{
							"msg":  "修改成功",
							"data": data,
							"code": "200",
						})
					}
				}
			}
		}
		db = db.Session(&gorm.Session{NewDB: true}) //必须清空上次遗留的链式条件
	})
}
