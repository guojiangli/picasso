package middleware

import (
	"github.com/valyala/bytebufferpool"
	"net/http"

	"github.com/gin-gonic/gin"
)

type okstatusWriter struct {
	gin.ResponseWriter
	// body
	body *bytebufferpool.ByteBuffer
	// header
	h http.Header

	wroteHeader bool
	code        int
}

func (tw *okstatusWriter) Write(b []byte) (int, error) {
	return tw.body.Write(b)
}

func (tw *okstatusWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	if tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *okstatusWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

func (tw *okstatusWriter) WriteHeaderNow() {}

func (tw *okstatusWriter) Header() http.Header {
	return tw.h
}

func OKStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		buffer := getBuff()
		tw := &okstatusWriter{body: buffer, ResponseWriter: c.Writer, h: make(http.Header)}
		c.Writer = tw
		c.Next()
		dst := tw.ResponseWriter.Header()
		for k, vv := range tw.Header() {
			dst[k] = vv
		}
		tw.ResponseWriter.WriteHeader(http.StatusOK)
		tw.ResponseWriter.Write(buffer.Bytes())
		putBuff(buffer)
	}
}
