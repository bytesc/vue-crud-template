package server

import (
	"github.com/gin-gonic/gin"
	"go_crud/mysql_db"
	"gorm.io/gorm"
	"strconv"
)

func PingGET(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "请求成功",
		})
	})
}

func QueryGET(r *gin.Engine, db *gorm.DB) {
	r.GET("api/user/list/:name", func(c *gin.Context) {
		name := c.Param("name")
		var dataList []mysql_db.CrudList
		db.Where("name = ?", name).Find(&dataList)
		if len(dataList) == 0 { //没有查到
			c.JSON(200, gin.H{
				"msg":  "查询失败，数据不存在",
				"data": dataList,
				"code": "400",
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功",
				"data": dataList,
				"code": "200",
			})
		}
	})
}

// QueryPageGET 分页查询
func QueryPageGET(r *gin.Engine, db *gorm.DB) {
	r.GET("api/user/list", func(c *gin.Context) {
		var dataList []mysql_db.CrudList
		var pageSize, pageNum int
		pageSizeStr := c.Query("pageSize")
		pageNumStr := c.Query("pageNum")
		if pageSizeStr != "" {
			pageSizeInt, err := strconv.Atoi(pageSizeStr) //?pageSize=xxx 查询参数
			pageSize = pageSizeInt
			// 127.0.0.1:8080/user/list/?pageNum=2&pageSize=3
			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "查询失败，pageNum参数格式错误",
					"data": err,
					"code": "400",
				})
			}
		} else {
			pageSize = -1
		}
		if pageNumStr != "" {
			pageNumInt, err := strconv.Atoi(pageNumStr)
			pageNum = pageNumInt
			if err != nil {
				c.JSON(200, gin.H{
					"msg":  "查询失败，pageSize参数格式错误",
					"data": err,
					"code": "400",
				})
			}
		} else {
			pageNum = 1
		}

		var total int64 //保存数据调数
		db.Model(dataList).Count(&total).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&dataList)
		//limit(-1)表示查询全部数据，offset表示跳过多少条

		if len(dataList) == 0 { //没有查到
			c.JSON(200, gin.H{
				"msg":  "查询失败，数据不存在",
				"data": dataList,
				"code": "400",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "成功",
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
				"code": "200",
			})
		}
	})
}
