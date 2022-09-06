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
	req := &apppb.CreateAppReq{}
	err := util.ReadJsonReq(req, c)
	if err != nil {
		util.AbortWithError(err, c)
		return
	}

	resp := &apppb.CreateAppResp{
		Error: "",
		Data: &apppb.CreateAppResp_AppContainer{
			App: req.App,
		},
	}

	util.WriteResp(resp, c)
}
