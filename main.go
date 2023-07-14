package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// 全局数据库
var (
	DB *gorm.DB
)

// Todo Module
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 初始化数据库
func initMySQL() (err error) {
	dsn := "root:011214@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping() //数据库是否能ping通
}

func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble;
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}

	defer DB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//告诉gin框架去哪里找
	r.LoadHTMLGlob("templates/*")
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static/css", "./static/css")
	r.Static("/static/js", "./static/js")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//v1
	v1group := r.Group("v1")
	{
		// 代办事项
		//添加
		v1group.POST("/todo", func(c *gin.Context) {
			//1.从请求中把数据拿出来
			var todo Todo
			c.ShouldBind(&todo)
			//2.存入数据库
			//3.返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//查看所有的代办事项
		v1group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			if err = DB.Find(&todos).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
			} else {
				c.JSON(http.StatusOK, todos)
			}
		})
		//修改
		v1group.PUT("/todo/:id", func(c *gin.Context) {
			//这里将状态修改交给前端处理，后端只要负责数据库中的状态更新和错误处理
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
				return
			}
			var todo Todo
			err = DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				todo.Status = true //对应数据库中的1
				DB.Save(&todo)
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除
		v1group.DELETE("/todo/:id", func(c *gin.Context) {
			//这里将删除交给前端处理，后端只要负责数据库中的删除更新和错误处理
			id := c.Param("id")
			var todo Todo
			if err = DB.Where("id=?", id).Delete(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run(":8080")
}
