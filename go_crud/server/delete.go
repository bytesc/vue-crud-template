package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
)

func DeletePOST(r *gin.Engine, db *gorm.DB) {
	r.POST("api/user/delete/:id", func(c *gin.Context) {
		var data []mysql_db.CrudList //切片类型，查询结果列表
		id := c.Param("id")          //接收路径参数
		// c.Query()  //接收查询参数
		db.Where("id = ?", id).Find(&data) //数据库查找
		if data == nil {                   //没有查到
			c.JSON(200, gin.H{
				"msg":  "删除失败，数据不存在",
				"data": data,
				"code": "400",
			})
		} else {
			result := db.Where("id = ?", id).Delete(&data) //删除对应数据
			if result.Error != nil {
				fmt.Println(result.Error)
				c.JSON(200, gin.H{
					"msg":  "删除失败",
					"data": result.Error,
					"code": "400",
				})
			} else {
				c.JSON(200, gin.H{
					"msg":  "删除成功",
					"data": data,
					"code": "200",
				})
			}
		}

	}) //delete请求也可
}
