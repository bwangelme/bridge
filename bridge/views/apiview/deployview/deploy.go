package deployview

import (
	"bridge/bridge/controller/deploycon"
	"bridge/bridge/log"
	"bridge/bridge/views/util"
	"github.com/gin-gonic/gin"
)

func GetDeploy(c *gin.Context) {
	log.L.Infof("get info of deploy %v", c.Param("id"))
	pods, err := deploycon.GetPods(c.Request.Context(), "rook-ceph")
	if err != nil {
		util.AbortWithCode(500, err, c)
		return
	}

	c.JSON(200, pods)
}
