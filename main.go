package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
	"net/http"
	"time"
)

// @title bubble接口文档
// @version 1.0
// @description bubble
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

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      routers.SetRouter(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("bubble start failed")
	}
}
