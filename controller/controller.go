package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"bubble/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1.获取请求数据
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.将数据存入数据库
	if err := models.CreateTodo(&todo); err != nil {
		// 3.返回响应
		c.JSON(http.StatusOK, gin.H {"err": err.Error()})
	}else{
		// 3.返回响应
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllTodo(c *gin.Context) {
	// 1.查询todos这个表，取出所有数据
	var todoList []models.Todo
	if err := models.GetAllTodo(&todoList); err != nil {
		// 2.返回数据
		c.JSON(http.StatusOK, gin.H {"msg": err.Error()})
	}else{
		// 2.返回数据
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H {"msg": err.Error()})
		return
	}
	c.BindJSON(&todo) // 将JSON参数绑定到todo
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H {"msg": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H {"msg": err})
	}else{
		c.JSON(http.StatusOK, gin.H {"msg": "ok"})
	}
}

