package appview

import "github.com/gin-gonic/gin"

func ListApp(c *gin.Context) {
	c.JSON(200, gin.H{
		"error": "",
		"apps":  nil,
	})
}
