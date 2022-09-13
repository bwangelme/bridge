package appview

import (
	"bridge/bridge/model/appmodel"
	"bridge/bridge/pbs/apppb"
	"bridge/bridge/views/util"

	"github.com/gin-gonic/gin"
)

func ListApp(c *gin.Context) {
	req := &apppb.ListAppReq{}
	err := util.ReadJsonReq(req, c)
	if err != nil {
		return
	}

	appIds, err := appmodel.List(req.Start, req.Limit)
	if err != nil {
		util.AbortWith500(err, c)
		return
	}

	apps, err := appmodel.GetMulti(appIds)
	if err != nil {
		util.AbortWith500(err, c)
		return
	}

	resp := &apppb.ListAppResp{
		Error: "",
		Data: &apppb.ListAppResp_AppContainer{
			Apps: appmodel.AppsToPB(apps),
		},
	}
	util.WriteResp(resp, c)
}

func CreateApp(c *gin.Context) {
	req := &apppb.CreateAppReq{}
	err := util.ReadJsonReq(req, c)
	if err != nil {
		return
	}

	app, err := appmodel.Add(req.App)
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
