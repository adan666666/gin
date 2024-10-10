package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(context *gin.Context) {
	accessToken := context.Request.Header.Get("access_token")
	if accessToken != token {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "身价认证失败",
		})
		context.Abort()
	}
	context.Next()
}

var token = "123456"

func Auth1(data interface{}) gin.HandlerFunc {
	fmt.Println(data)
	return func(context *gin.Context) {
		accessToken := context.Request.Header.Get("access_token") //请求头
		if accessToken != token {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "身价认证失败",
			})
			context.Abort() //终止请求
		}
		context.Set("access_token", accessToken) //可以在中间价中存值，跨中间介取值
		context.Next()                           //会执行下一个中间介，
	}
}
