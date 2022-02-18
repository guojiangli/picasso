package kfile

import (
	"fmt"
	"strings"
	"time"
)

// 从路径中获取文件名
// param: filepath 文件路径或者文件名
// notice: 如果不是路径 则认为是文件
func GetFileName(path string) string {
	if strings.Count(path, "/") == 0 && strings.Count(path, `\`) == 0 {
		return path
	}
	if strings.Count(path, "/") > strings.Count(path, `\`) {
		return path[strings.LastIndex(path, "/")+1:]
	}
	if strings.Count(path, "/") < strings.Count(path, `\`) {
		return path[strings.LastIndex(path, `\`)+1:]
	}
	return path
}

//  按年月日生成路径
//  notice: 这个只有阿里云上传的时候创建路径用 其他地方尽量不要使用
// 			最终类似:  2019/January/10
func GetDirectoryPathByDate() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	return fmt.Sprintf("%d/%s/%d", year, month, day)
}
