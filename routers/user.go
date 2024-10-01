package routers

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	//v1.Use() //在哪里use作用域就在哪个地方
	v1.GET("/user/:id/:name", controllers.GetUser)
	v1.POST("/user", controllers.AddUser)
	v1.PUT("/user", controllers.UpdateUser)
	v1.DELETE("/user", controllers.DeleteUser)

	v2 := group.Group("/v2")
	v2.GET("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "get user v2",
		})
	})
	v2.POST("/user")
	v2.PUT("/user")
	v2.DELETE("/user")
}

// 在main函数之前，导入包的时间 就会调用这个方法
func init() {

}
