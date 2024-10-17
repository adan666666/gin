package main

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql" //数据库驱动
)

type User struct {
	//匿名字段
	gorm.Model //相当于四个字段
	Age        int
	Name       string
	//Birthday time.Time
}
type Student struct {
	StuId int    `gorm:"primary_key;AUTO_INCREMENT"` //主键 自动增长的
	Name  string `gorm:"not null"`
	Age   int    `gorm:"index:name_index"` //索引
	Email string `gorm:"unique"`
	Sex   string
	Desc  string `gorm:"-"` //忽略
}

func main() {
	db, err := gorm.Open("mysql", "root:mima123@tcp(localhost:3306)/cg_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//数据库资源释放(用于确保数据库连接在函数执行完毕后被关闭)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	//开启打印日志
	db.LogMode(true)
	//User-users
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	//db.CreateTable(&Student{})
	//这种方式可以指定表名
	//db.Table("users").CreateTable(&User{})

	//删除表
	//db.DropTable(&User{})
	//db.DropTable("users")

	//判断表是否存在
	//flag1 := db.HasTable(&User{})
	//fmt.Println(flag1)
	//flag2 := db.HasTable("users")
	//fmt.Println(flag2)

	////增删改查
	//db.Create(&User{Age: 18, Name: "丽丽"}) //增加
	////查询
	//var mysuer User
	//db.First(&mysuer, "age = ?", 18)
	//fmt.Println(mysuer)
	//
	////更新数据
	//db.Model(&mysuer).Update("age", 30)
	//db.Model(&mysuer).Update("name", "菲菲")

	//删除数据
	//db.Delete(&mysuer)

	////支持原生sql
	//var users []User
	//db.Raw("select * from users age=?", 10).Find(&users)
	//fmt.Println(users)

}
