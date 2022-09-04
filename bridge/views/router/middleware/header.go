package middleware

import "github.com/gin-gonic/gin"

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Proma", "no-cache")
	c.Header("Expires", "0")
	c.Next()
}
