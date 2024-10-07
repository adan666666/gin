package stuope

import (
	"gin/dbope"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	//操作数据库
	//添加一条记录
	var student = models.Student{
		Name: "丽丽",
		Age:  18,
	}
	//添加操作
	dbope.DB.Create(&student)
	//关闭资源
	dbope.DB.Close()
}
