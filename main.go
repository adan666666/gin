package main

import (
	"fmt"
	"gin/middleware"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("This is a log message")
	r := gin.Default()
	api := r.Group("/api")
	api.Use(middleware.Auth1("like 12345"))
	fmt.Print("hhhhhhh")
	routers.InitUser(api)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
