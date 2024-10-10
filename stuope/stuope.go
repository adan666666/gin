package stuope

import (
	"encoding/json"
	"fmt"
	"gin/dbope"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Hello2(context *gin.Context) {
	var input struct {
		Username string `json:"username" gorm:""` //不写json也行。 但是字段名称会变成大写,对前端有影响
		Password string `json:"password"`
	}
	//接收json类型  ShouldBind绑定客户端传过来的json
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 将结构体序列化为JSON字符串
	jsonData, _ := json.Marshal(input)
	fmt.Println(string(jsonData))
	context.JSON(http.StatusOK, input) //直接传user也可以返回json
}
