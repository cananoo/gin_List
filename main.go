package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"list_Project/dao"
	"list_Project/models"
	"list_Project/routers"
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

	routers.SetupRouter().Run(":8080")
}
