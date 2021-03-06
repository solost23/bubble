package model

import (
	"bubble/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查

// CreateTodo 创建todo
func CreateTodo(todo *Todo) (err error) {
	err = mysql.DB.Create(&todo).Error
	return err
}

// GetAllTodo 查找所有todo
func GetAllTodo(todoList *[]Todo) (err error) {
	err = mysql.DB.Find(&todoList).Error
	return err
}

// GetATodo 查找一个todo
func GetATodo(id string) (todo *Todo, err error) {
	if err = mysql.DB.Where("id=?", id).First(todo).Error; err != nil {

		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = mysql.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = mysql.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
