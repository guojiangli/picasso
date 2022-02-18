package middleware

import (
	"github.com/valyala/bytebufferpool"
)

const BuffSize = 10 * 1024

//var buffPool sync.Pool

func getBuff() *bytebufferpool.ByteBuffer {
	buffer:=bytebufferpool.Get()
	return buffer
}

func putBuff(buffer *bytebufferpool.ByteBuffer) {
	buffer.Reset()
	bytebufferpool.Put(buffer)
}
