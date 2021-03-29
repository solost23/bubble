package main

import (
	"fmt"
	"github.com/go_web/lesson25/dao"
	"github.com/go_web/lesson25/models"
	"github.com/go_web/lesson25/routers"

)

func main() {
	// 创建数据库 CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	defer dao.DB.Close()

	// 创建数据表
	dao.DB.AutoMigrate(&models.Todo{})

	router := routers.SetRouter()

	router.Run(":8080")
}
