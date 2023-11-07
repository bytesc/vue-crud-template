package main

import (
	"fmt"
	"go_crud/mysql_db"
	"go_crud/server"
)

func main() {

	db, err := mysql_db.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 数据库迁移
	//err = db.AutoMigrate(&mysql_db.CrudList{})

	r := server.CreateServer()

	server.PingGET(r)

	server.AddPOST(r, db)

	r.Run("0.0.0.0:8080") // 监听并在 0.0.0.0:8080 上启动服务
	// http://127.0.0.1:8080/ping
	//fmt.Println(r)

}
