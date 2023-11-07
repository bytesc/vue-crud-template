package server

//package main

import (
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()
	return r
}

//r.Run("0.0.0.0:8080") // 监听并在 0.0.0.0:8080 上启动服务
//// http://127.0.0.1:8080/ping
////fmt.Println(r)

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
