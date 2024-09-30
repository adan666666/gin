package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	id := c.DefaultQuery(`id`, "0")
	name := c.DefaultQuery(`name`, "0")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": name,
	})
}
func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	addr := c.PostForm("addr")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
		"addr": addr,
	})
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
