package main

import (
	"fmt"
	"go-crud/connect-db"
	"go-crud/server"
)

func main() {

	db, err := connect_db.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 数据库迁移
	err = db.AutoMigrate(&connect_db.CrudList{})

	r := server.StartServer()
	fmt.Println(r)

}
