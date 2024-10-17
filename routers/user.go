package routers

import (
	"gin/controllers"
	"gin/stuope"
	"github.com/gin-gonic/gin"
)

/**
总路由
*/

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	//v1.Use(middleware.Auth1("like 666")) //在哪里use作用域就在哪个地方=在v1的路由上
	v1.GET("/user/:id/:name", controllers.GetUser)
	v1.POST("/user", controllers.AddUser)
	v1.PUT("/user", controllers.UpdateUser) //put
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

	//分路由
	s := group.Group("/stuope") //添加Student操作分组
	//模块分组
	stuope.Router(s)

}

// 在main函数之前，导入包的时间 就会调用这个方法
//func init() {
//
//}
