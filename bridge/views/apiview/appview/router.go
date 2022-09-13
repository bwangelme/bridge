package appview

import "github.com/gin-gonic/gin"

func InitRouter(r gin.IRouter) {
	r.GET("/app", ListApp)
	r.POST("/app", CreateApp)
}
