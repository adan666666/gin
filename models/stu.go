package models

type Student struct {
	StuId int    `gorm:"primary_key;AUTO_INCREMENT"` //主键 自动增长的
	Name  string `gorm:"not null"`
	Age   int    `gorm:"index:name_index"` //索引
	Email string `gorm:"unique"`
	Sex   string
	Desc  string `gorm:"-"` //忽略
}
