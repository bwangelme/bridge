package middleware

import (
	"bridge/bridge/log"
	"bridge/bridge/views/httperr"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandlerHttpErr(c *gin.Context) {
	c.Next()

	length := len(c.Errors)
	if c.IsAborted() && length > 0 {
		e := c.Errors[length-1]
		var hErr *httperr.HTTPErr
		ok := errors.As(e, &hErr)
		if ok {
			c.JSON(hErr.Code, gin.H{
				"err":  hErr.Err.Error(),
				"data": nil,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":  "server internal error",
				"data": nil,
			})
			log.L.WithFields(logrus.Fields{
				"msg": "unknown error",
				"err": e,
			}).Error()
			return
		}
	}
}
