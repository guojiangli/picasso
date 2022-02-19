package middleware

import (
	"strconv"
	"time"

	"github.com/guojiangli/picasso/pkg/kmetric"
	"github.com/gin-gonic/gin"
)

func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		beg := time.Now()
		c.Next()
		kmetric.ServerHandleHistogram.Observe(time.Since(beg).Seconds(), kmetric.TypeHTTP, c.Request.URL.Path, c.Request.Method)
		kmetric.ServerHandleCounter.Inc(kmetric.TypeHTTP, c.Request.URL.Path, c.Request.Method, strconv.Itoa(c.Writer.Status()))
		return
	}
}
