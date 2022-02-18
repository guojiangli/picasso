package kfile

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

// 获取系统换行符号
// notice: 其他系统自行兼容吧
func GetSystemLineBreak() (lineBreak string) {
	if runtime.GOOS == "windows" {
		lineBreak = "\r\n"
	} else {
		lineBreak = "\n"
	}
	return
}

// 获取系统分隔符
// notice: 其他系统自行兼容吧
func GetSystemDelimiter() (delimiter string) {
	if runtime.GOOS == "windows" {
		delimiter = `\`
	} else {
		delimiter = "/"
	}
	return
}

// 获取文件大小
// param: filepath 文件路径或者文件名
func GetFileSize(filepath string) (int64, error) {
	if info, err := os.Stat(filepath); err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

// 从路径中读取文件名与后缀
// param: filepath 文件路径或者文件名
func GetSuffixAndFilename(filepath string) (filename, suffix string) {
	if strings.Count(filepath, ".") > 0 {
		suffix = filepath[strings.LastIndex(filepath, ".")+1:]
	}
	var delimiter string
	if runtime.GOOS == "windows" {
		delimiter = "\\"
	} else {
		delimiter = "/"
	}
	if strings.Count(filepath, delimiter) > 0 {
		filename = filepath[strings.LastIndex(filepath, "/")+1:]
	} else {
		filename = filepath
	}
	return
}

// 检查目录中是否有指定文件
// param: filename 文件名 可以是绝对路径
// param: localPath  本地文件夹
// notice: 如果读取目录失败,那么将会认为目录中没有这个文件
func CheckFileInDisk(filename, localPath string) bool {
	files, dirErr := ioutil.ReadDir(localPath)
	if dirErr != nil {
		return false
	}
	for _, value := range files {
		if value.Name() == GetFileName(filename) {
			return true
		}
	}
	return false
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(input)), nil
}
