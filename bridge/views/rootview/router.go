package rootview

import "github.com/gin-gonic/gin"

func InitRouter(r gin.IRouter) {
	root := r.Group("/")

	root.GET("/", Index)
	root.GET("/ping", Ping)
}
