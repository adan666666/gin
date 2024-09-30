package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//http://localhost:8080/user?name=张三&age=100
	r.GET("/user", func(c *gin.Context) {
		id := c.DefaultQuery(`id`, "0")
		name := c.DefaultQuery(`name`, "0")
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": name,
		})
	})
	//get通过路径传参 ：http://localhost:8080/user1/10/张三
	r.GET("/user1/:id/:name", func(c *gin.Context) {
		id := c.Param("id") //路径传这里要使用Param
		name := c.Param(`name`)
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"message": name,
		})
	})
	//PostForm 是form表单：使用x-www-form-urlencoded这个传值
	r.POST("/user", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		addr := c.PostForm("addr")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"addr": addr,
		})
	})
	r.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "put",
		})
	})
	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
