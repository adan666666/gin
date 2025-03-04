package main

import (
	"fmt"
	"gin/config"
	"gin/middleware"
	"gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main函数")
	fmt.Println("你好")
	config.InitConfig()
	fmt.Println("yml配置信息", config.AppConfig.Database.Dsn)
	r := gin.Default()
	api := r.Group("/api")
	api.Use(middleware.Auth1("like 12345"))
	fmt.Println("hhhhhhh")
	routers.InitUser(api)
	//端口
	r.Run("localhost:8080")

}
