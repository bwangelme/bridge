package appview

import (
	"bridge/bridge/pbs/apppb"
	"bridge/bridge/views/util"

	"github.com/gin-gonic/gin"
)

func ListApp(c *gin.Context) {
	c.JSON(200, gin.H{
		"error": "",
		"data":  nil,
	})
}

func CreateApp(c *gin.Context) {
	app := &apppb.App{
		Name:    "bridge",
		Runtime: apppb.App_Golang,
	}

	resp := &apppb.CreateAppResp{
		Error: "",
		Data: &apppb.CreateAppResp_AppContainer{
			App: app,
		},
	}

	util.WriteResp(resp, c)
}
