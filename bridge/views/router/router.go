package router

import (
	"bridge/bridge/views/appview"
	"bridge/bridge/views/commonview"
	"bridge/bridge/views/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(middlewares ...gin.HandlerFunc) http.Handler {
	r := gin.New()
	r.Use(middlewares...)

	r.StaticFile("/favicon.ico", "src/static/img/favicon.ico")
	r.Static("/static", "src/static")
	r.LoadHTMLGlob("templates/**")

	r.GET("/", commonview.Index)
	r.GET("/ping", commonview.Ping)

	api := r.Group("/api")
	api.Use(middleware.NoCache)

	{
		api.GET("/app", appview.ListApp)
	}

	return r
}
