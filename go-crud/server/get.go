package server

import "github.com/gin-gonic/gin"

func PingGET(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "请求成功",
		})
	})
}
