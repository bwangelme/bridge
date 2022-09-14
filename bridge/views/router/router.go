package router

import (
	"bridge/bridge/views/apiview"
	"bridge/bridge/views/rootview"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(middlewares ...gin.HandlerFunc) http.Handler {
	r := gin.New()
	r.Use(middlewares...)

	r.StaticFile("/favicon.ico", "src/static/img/favicon.ico")
	r.Static("/static", "src/static")
	r.LoadHTMLGlob("templates/**")

	rootview.InitRouter(r)
	apiview.InitRouter(r)

	return r
}
