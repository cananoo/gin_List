package models

import (
	"list_Project/dao"
)

// Todo Module
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/**
  Todo 增删改查操作
*/

func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}
func GetTodoList() (todos []*Todo, err error) {

	if err = dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	ntodo := new(Todo)
	if err = dao.DB.Where("id=?", id).First(ntodo).Error; err != nil {
		return nil, err
	}
	return ntodo, nil
}

func UpdateATodo(todo *Todo) (err error) {
	todo.Status = true //对应数据库中的1
	if err = dao.DB.Save(todo).Error; err != nil {
		return
	}
	return

}

func DeleteATodo(id string) (err error) {
	if err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
		return
	}
	return
}
