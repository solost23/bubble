package main

import (
	"net/http"
	"time"

	"bubble/config"
	"bubble/model"
	"bubble/mysql"
	"bubble/routers"
)

// @title bubble接口文档
// @version 1.0
// @description bubble
// @Schemes http https
func main() {
	// 创建数据库 CREATE DATABASE bubble;
	// 连接数据库
	err := mysql.InitDB()
	if err != nil {
		panic(err)
	}
	defer mysql.DB.Close()

	// 创建数据表
	mysql.DB.AutoMigrate(&model.Todo{})

	serviceConfig := config.NewProject()
	server := &http.Server{
		Addr:         serviceConfig.ServiceAddr + ":" + serviceConfig.ServicePort,
		Handler:      routers.SetRouter(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("bubble start failed")
	}
}
