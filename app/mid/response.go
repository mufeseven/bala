package mid

import (
	"bala/app/log"
	"bytes"

	"github.com/gin-gonic/gin"
)

type responseBodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseBodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LogResponse(c *gin.Context) {
	blw := &responseBodyLogWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
	c.Writer = blw
	c.Next()
	log.Local().Info(blw.body.String())
	blw.body.Reset()
}
