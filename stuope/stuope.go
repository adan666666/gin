package stuope

import (
	"encoding/json"
	"fmt"
	"gin/dbope"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	//dbope.DB.Close()
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

// 分页
func StudentList(context *gin.Context) {

	var dataList []models.Student
	pageSize, _ := strconv.Atoi(context.Query("pageSize")) ///* strconv.Atoi 字符串转int*/
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}
	fmt.Println("测试：", pageSize)
	fmt.Println("测试：", pageNum)
	//返回一个总数
	var total int64
	dbope.DB.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList) //dbope.DB.Model(&dataList)//这里不能传地址要不然查询的是空的
	fmt.Println("测试 total=", total)

	if len(dataList) == 0 {
		context.JSON(http.StatusOK,
			gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{}})
	} else {
		context.JSON(
			http.StatusOK,
			gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageSize": pageSize,
					"pageNum":  pageNum,
				},
			},
		)
	}
	//关闭资源
	//dbope.DB.Close()
}
