package util

import (
	"bridge/bridge/views/httperr"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ReadJsonReq(req proto.Message, c *gin.Context) error {
	jsonData, err := c.GetRawData()
	if err != nil {
		return httperr.New400Error(errors.WithMessagef(err, "get raw data failed"))
	}
	err = protojson.Unmarshal(jsonData, req)
	if err != nil {
		return httperr.New400Error(errors.WithMessagef(err, "unmarshal failed"))
	}

	return nil
}

func AbortWithError(err error, c *gin.Context) {
	// ignore the gin.parsedError
	_ = c.Error(err)
	c.Abort()
}

func WriteResp(resp proto.Message, c *gin.Context) {
	res, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(resp)
	if err != nil {
		err := httperr.New500Error(errors.WithMessagef(err, "marshal json failed"))
		AbortWithError(err, c)
		return
	}

	c.Data(200, gin.MIMEJSON, res)
}
