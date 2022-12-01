package common

import (
	"fmt"
	"go_gin/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"

	username := "root"
	password := "123456"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	fmt.Println(args)
	//args := "root:12345678@127.0.0.1:3306/ginessential?charset=utf8&parseTime=true"
	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("数据库连接错误:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
