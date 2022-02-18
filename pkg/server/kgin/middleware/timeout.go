package middleware

import (
	"context"
	"github.com/valyala/bytebufferpool"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type timeoutWriter struct {
	gin.ResponseWriter
	// body
	body *bytebufferpool.ByteBuffer
	// header
	h http.Header

	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

func (tw *timeoutWriter) Write(b []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}

	return tw.body.Write(b)
}

func (tw *timeoutWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

func (tw *timeoutWriter) WriteHeaderNow() {}

func (tw *timeoutWriter) Header() http.Header {
	return tw.h
}

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// 这个中间件包含mutex的加解锁操作,会导致G暂停,建议仍旧使用ContextTimeout.
func Timeout(t time.Duration, PanicResBody []byte, TimeoutResBody []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// sync.Pool
		buffer := getBuff()

		tw := &timeoutWriter{body: buffer, ResponseWriter: c.Writer, h: make(http.Header)}
		c.Writer = tw

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		// Channel capacity must be greater than 0.
		// Otherwise, if the parent coroutine quit due to timeout,
		// the child coroutine may never be able to quit.
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-panicChan:
			c.Abort()
			tw.ResponseWriter.WriteHeader(http.StatusInternalServerError)
			tw.ResponseWriter.Write(PanicResBody)

		case <-ctx.Done():
			tw.mu.Lock()
			defer tw.mu.Unlock()
			tw.ResponseWriter.WriteHeader(http.StatusServiceUnavailable)
			tw.ResponseWriter.Write(TimeoutResBody)
			c.Abort()
			tw.timedOut = true
			// If timeout happen, the buffer cannot be cleared actively,
			// but wait for the GC to recycle.
		case <-finish:
			tw.mu.Lock()
			defer tw.mu.Unlock()
			dst := tw.ResponseWriter.Header()
			for k, vv := range tw.Header() {
				dst[k] = vv
			}
			tw.ResponseWriter.WriteHeader(tw.code)
			tw.ResponseWriter.Write(buffer.Bytes())
			putBuff(buffer)
		}
	}
}
