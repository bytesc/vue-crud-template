package server

//package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/connect_db"
	"gorm.io/gorm"
)

func CreateServer() *gin.Engine {
	r := gin.Default()
	return r
}

func PingGET(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "请求成功",
		})
	})
}

func AddPOST(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/add", func(c *gin.Context) {
		var listRes connect_db.CrudList
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

func init() {

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "请求成功",
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// http://127.0.0.1:8080/ping
}
