package main

import (
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql" //数据库驱动
)

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

	//一对一
	//创建表：通常情况下，数据库中新建的标的名字是结构体名字的复数形式，例如结构体User，表名 users
	//db.CreateTable(&User{})
	//db.CreateTable(&UserInfo{})

	//var userinfo UserInfo
	//db.Preload("User").Find(&userinfo, &userinfo, "info_id=?", 1)
	//fmt.Printf("%#v\n", userinfo)

	//一对多
	//添加
	/*author := Author{
		Name: "张三",
		Age:  30,
		Sex:  "男",
		Article: []Article{
			{
				Title:   "HTML入门",
				Content: "HTML******",
				Desc:    "好的不得了",
			},
			{
				Title:   "CSS入门",
				Content: "CSS******",
				Desc:    "好的不得了2",
			},
		},
	}
	db.Create(&author)*/

	////关联查询
	//var author []Author
	//db.Preload("Article").Find(&author)
	//fmt.Println(author)

	//关联更新
	//1.先查询
	//2.再更新
	//db.Model(&author[0].Article).Where("ar_id=?", 1).Update("title", "JS入门--")

	//关联删除
	//1.先查询
	//2.再删除
	//db.Delete(&author[0].Article) //这样的话，会删除这个作用的所有的文章
	//db.Where("ar_id=?", 2).Delete(&author[0].Article)

	//db.Find(&user, []int{1, 2})
	//fmt.Printf("%+v", user)

	//user1 := User{·
	//	Age:  14,
	//	Name: "莎莎",
	//}
	//db.Save(&user1) //create和save是一样，都可以进行添加操作

}

type User struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"`
	Age    int
	Name   string
}
type UserInfo struct {
	InfoID  int `gorm:"primary_key;AUTO_INCREMENT"`
	Pic     string
	Address string
	Email   string
	//关联关系
	User User `gorm:"ForeignKey:UserId;AssociationForeignKey:UserId"`
	//指定外键
	UserId int
}

// 一对多
type Author struct {
	AID  int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Age  int
	Sex  string
	//关联关系：
	Article []Article `gorm:"ForeignKey:AuId;AssociationForeignKey:AID"`
}
type Article struct {
	ArId    int `gorm:"primary_key;AUTO_INCREMENT"`
	Title   string
	Content string
	Desc    string
	//设置外键：
	AuId int
}
