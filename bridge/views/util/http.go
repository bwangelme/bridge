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
		e := errors.WithMessagef(err, "get raw data failed")
		AbortWithCode(400, e, c)
		return e
	}
	err = protojson.Unmarshal(jsonData, req)
	if err != nil {
		e := errors.WithMessagef(err, "unmarshal failed")
		AbortWithCode(400, e, c)
		return e
	}

	return nil
}

func AbortWithCode(code int, err error, c *gin.Context) {
	e := httperr.NewHTTPErr(code, err)
	abortWithError(e, c)
}

func AbortWith500(err error, c *gin.Context) {
	AbortWithCode(500, err, c)
}

func abortWithError(err error, c *gin.Context) {
	// ignore the gin.parsedError
	_ = c.Error(err)
	c.Abort()
}

func WriteResp(resp proto.Message, c *gin.Context) {
	res, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(resp)
	if err != nil {
		err := httperr.New500Error(errors.WithMessagef(err, "marshal json failed"))
		abortWithError(err, c)
		return
	}

	c.Data(200, gin.MIMEJSON, res)
}
