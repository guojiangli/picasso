package kstring

import (
	"fmt"
	"time"
)

// Formatter ...
type Formatter string

// Format ...
func (fm Formatter) Format(args ...interface{}) string {
	return fmt.Sprintf(string(fm), args...)
}

func KVInterface(key string, value interface{}) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func KVString(key string, value string) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func KVInt(key string, value int) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func KVInt64(key string, value int64) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func KVInt32(key string, value int32) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func KVTime(key string, value time.Time) string {
	return fmt.Sprint(key) + ":{" + fmt.Sprintf("%+v", value) + "}"
}
func Title(title string) string {
	return "[" + fmt.Sprint(title) + "]"
}

func SugarFormat(vals ...interface{}) string {
	var dst string
	for i := 0; i < len(vals); i++ {
		if i > 0 {
			dst += " "
		}
		dst += fmt.Sprint(vals[i])
	}
	return dst
}
