package helper

import (
	"bridge/bridge/views/httperr"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ReadJsonReq(req proto.Message, c *gin.Context) (err error) {
	var jsonData []byte

	if c.Request.Method == http.MethodGet {
		var res = map[string]string{}
		for key, _ := range c.Request.URL.Query() {
			v, ok := c.GetQuery(key)
			if !ok {
				continue
			}
			res[key] = v
		}
		jsonData, err = json.Marshal(res)
		if err != nil {
			e := errors.WithMessagef(err, "marshal get form failed")
			AbortWithCode(400, e, c)
			return e
		}
	} else {
		jsonData, err = c.GetRawData()
		if err != nil {
			e := errors.WithMessagef(err, "get raw data failed")
			AbortWithCode(400, e, c)
			return e
		}
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
