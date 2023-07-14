package routers

import (
	"github.com/gin-gonic/gin"
	"list_Project/controller"
)

func SetupRouter() *gin.Engine {
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
	return r
}
