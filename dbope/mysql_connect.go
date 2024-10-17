package dbope

import (
	"fmt"
	"gin/models"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql" //数据库驱动
)

var DB *gorm.DB
var ERR error

func init() {
	fmt.Println("init函数")
	DB, ERR = gorm.Open("mysql", "root:mima123@tcp(localhost:3306)/cg_dev?charset=utf8&parseTime=True&loc=Local")
	if ERR != nil {
		panic(ERR)
	}
	//开启打印日志
	DB.LogMode(true)
	DB.CreateTable(&models.Student{})
}
