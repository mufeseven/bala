package mid

import (
	"bala/app/log"
	"bala/app/protocol"
	"bala/app/protocol/protocode"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EncodeRequestBody(code protocode.ProtoCode, v interface{}) (*protocol.RequestBody, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return &protocol.RequestBody{Protocol: code.String(), Packet: buf.String()}, nil
}

func DecodeRequestBody(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		body := &protocol.RequestBody{}
		err := c.BindJSON(body)
		if err != nil {
			log.Local().Errorf("DecodeRequestBody error:%s", err)
			c.Abort()
			return
		}
	}
}
