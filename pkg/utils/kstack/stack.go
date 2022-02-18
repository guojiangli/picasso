package kstack

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"

	jsoniter "github.com/json-iterator/go"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

type StackData struct {
	File     string  `json:"file"`
	Line     int     `json:"line"`
	Pc       uintptr `json:"pc"`
	Function string  `json:"function"`
	Source   string  `json:"source"`
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func Stack(skip int) []byte {
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	var datas []StackData
	for i := skip; ; i++ { // Skip the expected number of frames
		var d StackData
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// // Print this much at least.  If we can't find the source, it won't show.
		// fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		d.File = file
		d.Line = line
		d.Pc = pc
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				datas = append(datas, d)
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		d.Function = string(function(pc))
		d.Source = string(source(lines, line))
		datas = append(datas, d)
	}
	res, err := jsoniter.Marshal(datas)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
