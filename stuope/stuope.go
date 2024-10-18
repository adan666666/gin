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

// put 改
func Update(context *gin.Context) {
	var student models.Student
	stu_id := context.Param("stu_id") //http://localhost:8080/api/stuope/studentUpdate/1
	fmt.Println("测试=", stu_id)
	//断断id是否存在
	//dbope.DB.Where("id=?", ctx.Param("id")).Find(&data)
	//或者    不要Select默认选择所有
	dbope.DB. /*.Select("stu_id,name,age")*/ Where("stu_id = ?", stu_id).Find(&student) //// 获取所有匹配记录
	fmt.Println(student)
	//判断ID是否存在
	if student.StuId == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "没有查询到数据",
		})
		return
	}
	var s_id = student.StuId

	//绑定  //传过来的是json  //会重写上面student的值
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "错误请求参数"})
		return
	}
	//优化
	if student.StuId != s_id {
		student.StuId = s_id
	}
	fmt.Println(student) //会重写上面student的值

	//修改 绑定的student里面有id所以会根据id来更新
	//updates := dbope.DB.Model(&student). /*Where("stu_id = ?", stu_id).*/ Updates(&student)
	//这种student1里面没有id   客户端也不要传id过来  因为第一个student通过 .Param("stu_id")浏览器传过来查出来的，第二个是传过的，stu_id假如不一致会出问题
	//UPDATE `students` SET `age` = 18, `email` = 'xxxxx@gamil.com', `name` = '张三', `sex` = '女', `stu_id` = 7  WHERE `students`.`stu_id` = 7 and `students`.`stu_id` = 1
	//{
	//    "Name": "张三经",
	//    "Age": 18,
	//    "Email": "xxxxx@gamil.com",
	//    "Sex": "女",
	//    "Desc": ""
	//}
	updates := dbope.DB.Model(&student).Where("stu_id = ?", stu_id).Updates(&student)
	if updates.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新成功",
			"data": updates,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "更新失败",
		})
	}
}
