package appview

import (
	"bridge/bridge/model/appmodel"
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
		return
	}

	app, err := appmodel.AddApp(req.App)
	if err != nil {
		util.AbortWith500(err, c)
		return
	}

	resp := &apppb.CreateAppResp{
		Error: "",
		Data: &apppb.CreateAppResp_AppContainer{
			App: app.ToPB(),
		},
	}

	util.WriteResp(resp, c)
}
