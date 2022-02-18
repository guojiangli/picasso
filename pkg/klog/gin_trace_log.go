package klog

import "github.com/gin-gonic/gin"

func HTTPServerTraceLog(l *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(WithTraceCtx(l, c.Request.Context()))
		c.Next()
	}
}
