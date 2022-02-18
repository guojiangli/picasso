package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/bytebufferpool"
	"io/ioutil"
	"time"

	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/klog/baselogger"

	"github.com/gin-gonic/gin"
)

type logData struct {
	Method          string              `json:"method"`
	URI             string              `json:"uri"`
	Proto           string              `json:"proto,omitempty"`
	ClientIP        string              `json:"client_ip"`
	StatusCode      int                 `json:"status_code"`
	Latency         int64               `json:"latency(μs)"`
	RequestHeaders  map[string][]string `json:"request_headers,omitempty"`
	RequestBody     string              `json:"request_body,omitempty"`
	ResponseHeaders map[string][]string `json:"response_headers,omitempty"`
	ResponseBody    string              `json:"response_body,omitempty"`
	ErrorMessage    string              `json:"error_message"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytebufferpool.ByteBuffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func TotalLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := klog.FromTraceCtx(c.Request.Context())
		buffer := getBuff()
		bodyLogWriter := &bodyLogWriter{body: buffer, ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// Start timer
		r := c.Request
		start := time.Now()
		// Process request
		var bodyBytes []byte
		var err error
		if r.Body != nil {
			bodyBytes, err = ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Log("kgin.middleware.totallog error:", err)
			}
		}
		defer r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()

		Latency := time.Since(start).Microseconds()

		ErrorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		responseBody := bodyLogWriter.body.String()

		Data := logData{
			Method:          r.Method,
			URI:             r.RequestURI,
			Proto:           r.Proto,
			ClientIP:        c.ClientIP(),
			StatusCode:      c.Writer.Status(),
			Latency:         Latency,
			RequestHeaders:  r.Header,
			RequestBody:     string(bodyBytes),
			ResponseHeaders: c.Writer.Header(),
			ResponseBody:    responseBody,
			ErrorMessage:    ErrorMessage,
		}
		dataBytes, err := json.Marshal(Data)
		if err != nil {
			logger.Log("kgin.middleware.totallog error:", err)
		}
		logger.Log(string(dataBytes))
		putBuff(buffer)
	}
}

func SimpleLog(logger baselogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		start := time.Now()
		// Process request
		c.Next()

		Latency := time.Since(start).Microseconds()

		ErrorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		Data := logData{
			Method:       r.Method,
			URI:          r.RequestURI,
			ClientIP:     c.ClientIP(),
			StatusCode:   c.Writer.Status(),
			Latency:      Latency,
			ErrorMessage: ErrorMessage,
		}
		dataBytes, _ := json.Marshal(Data)
		logger.Log(string(dataBytes))
	}
}
