package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)


func SetRouter() *gin.Engine{
	router := gin.Default()

	// 网页文件解析
	router.LoadHTMLGlob("./templates/*")
	// 静态文件解析
	router.Static("/static", "./static")

	router.GET("/", controller.IndexHandler)

	// 待办事项
	// v1
	v1Group := router.Group("/v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看
		// 1.查看所有待办事项
		v1Group.GET("/todo", controller.GetAllTodo)
		// 修改某一个待办事项
		v1Group.PUT("todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("todo/:id", controller.DeleteATodo)
	}
	return router
}

