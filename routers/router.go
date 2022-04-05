package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"bubble/controller"
	_ "bubble/docs"
)

func SetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
