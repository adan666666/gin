package stuope

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("tostu", Hello1)
	r.POST("testJson", Hello2)
}
