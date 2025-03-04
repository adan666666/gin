package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 从Gin中发布的JSON获取特定参数的方法如下：
// 1. 使用Gin的Context.BindJSON()方法将JSON数据绑定到一个struct中，struct中的字段名称必须与JSON中的参数名称相同；
// 2. 使用Gin的Context.Get()方法从struct中获取特定参数；
// 3. 使用Gin的Context.Query()方法从URL中获取特定参数；
// 4. 使用Gin的Context.PostForm()方法从表单中获取特定参数；
// 5. 使用Gin的Context.Params()方法从路由参数中获取特定参数,和Query一样
func GetUser(c *gin.Context) {
	id := c.DefaultQuery(`id`, "0")
	name := c.DefaultQuery(`name`, "0")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": name,
	})
}
func AddUser(c *gin.Context) {
	req := &user{}
	//form表单和json都可以绑定
	err := c.ShouldBind(req) //Bind发生错误会直接返回给客户，程序不往下走
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	j, _ := json.Marshal(req) //结构体转json
	fmt.Println(string(j))
	c.JSON(http.StatusOK, req)
	/*name := c.PostForm("name")
	age := c.PostForm("age")
	addr := c.PostForm("addr")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
		"addr": addr,
	})*/
}
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "update user",
	})
}
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete users",
	})
}

// 在go语言里首字母大写是公有的，小写是私有的
// form(表单标签)、binding(表单验证标签)
// form:"name"  相当于定一个别名前端传过来的字段
type user struct {
	//可同时加增加json和form标签
	Name  string `json:"name" form:"name" binding:"required,alphaunicode"` //github.com/go-playground/validator/v10  这个库的验证
	Age   int    `form:"age" binding:"number"`
	Addr  string `form:"addr" binding:"alphaunicode"`
	Email string `form:"email" binding:"omitempty,email"` //如果为不为空再校验
}
