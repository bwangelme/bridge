package apiview

import (
	"bridge/bridge/views/apiview/appview"
	"bridge/bridge/views/router/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r gin.IRouter) {
	api := r.Group("/api")

	api.Use(middleware.NoCache)
	api.Use(middleware.HandlerHttpErr)

	// init sub router
	appview.InitRouter(api)
}
