package deployview

import "github.com/gin-gonic/gin"

func InitRouter(r gin.IRouter) {
	r.GET("/deploy/:id", GetDeploy)
}
