package dao

import "github.com/jinzhu/gorm"

// 全局数据库
var (
	DB *gorm.DB
)

// 初始化数据库
func InitMySQL() (err error) {
	dsn := "root:011214@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping() //数据库是否能ping通
}

func Close() {
	DB.Close()
}
