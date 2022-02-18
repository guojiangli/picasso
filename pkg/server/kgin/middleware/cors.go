package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		//跨域自定义header ，不设置自动加入，在下面手动设置
		cosHeader := c.GetHeader("Access-Control-Request-Headers")
		cosHeaders := strings.Split(cosHeader, ",")
		for _, v := range cosHeaders {
			headerKeys = append(headerKeys, v)
		}

		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		//headerStr += ",user"
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", headerStr)

			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "POST, GET,option, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,seed,sign")
			c.Header("Access-Control-Allow-Origin", origin)
			c.Set("content-type", "application/json")
		}

		//放行所有option方法
		if strings.ToUpper(method) == "option" {
			c.JSON(http.StatusOK, "option Request!")
		}
		c.Header("Access-Control-Max-Age", "600")
		c.Next()

	}
}
