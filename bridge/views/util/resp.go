package util

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteResp(resp proto.Message, c *gin.Context) {
	res, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(resp)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
			"data":  nil,
		})
		return
	}

	c.Data(200, gin.MIMEJSON, res)
}
