package controller

import (
	"github.com/gin-gonic/gin"
	"list_Project/models"
	"net/http"
)

/*
*
url    --> controller -->logic   -->model
请求来了 --> 控制器      -->业务逻辑 -->模型层的增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

func CreateATodo(c *gin.Context) {
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.ShouldBind(&todo)
	//2.存入数据库
	//3.返回响应
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func GetTodoList(c *gin.Context) {

	if todos, err := models.GetTodoList(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func UpdateATodo(c *gin.Context) {
	//这里将状态修改交给前端处理，后端只要负责数据库中的状态更新和错误处理
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		models.UpdateATodo(todo)
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	//这里将删除交给前端处理，后端只要负责数据库中的删除更新和错误处理
	id := c.Param("id")
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
