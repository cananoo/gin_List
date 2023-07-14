package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"list_Project/controller"
	"list_Project/dao"
	"list_Project/models"
)

func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	defer dao.Close() //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()
	//告诉gin框架去哪里找
	r.LoadHTMLGlob("templates/*")
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static/css", "./static/css")
	r.Static("/static/js", "./static/js")
	r.GET("/", controller.IndexHandler)
	//v1
	v1group := r.Group("v1")
	{
		// 代办事项
		//添加
		v1group.POST("/todo", controller.CreateATodo)
		//查看所有的代办事项
		v1group.GET("/todo", controller.GetTodoList)
		//修改
		v1group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	r.Run(":8080")
}
