package test

import (
	"bala/app"
	"bala/app/config"
	"bala/app/data"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Get(url string, args ...interface{}) *httptest.ResponseRecorder {
	// 构造get请求
	req := httptest.NewRequest("GET", fmt.Sprintf(url, args...), nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应的handler接口
	engine.ServeHTTP(w, req)
	return w
}

func PostForm(url string, body interface{}, args ...interface{}) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", fmt.Sprintf(url, args...), toReader(body))
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	engine.ServeHTTP(w, req)
	return w
}

func toReader(v interface{}) io.Reader {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(v)
	return buf
}

func init() {
	config.GoTest = true
	config.Path = "../../config"
	data.Path = "../../data"
	injector, _, err := app.BuildInjector()
	if err != nil {
		panic(err)
	}
	engine = injector.Engine
}
