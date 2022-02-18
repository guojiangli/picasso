package kfile

import "fmt"

func TransCodeSize(size int64) string {
	sizeF := float64(size)
	sizeMap := map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
	}
	sizeStr := ""
	for i := 0; i < 5; i++ {
		if sizeF < 10 {
			break
		}
		sizeF = sizeF / 1024
		sizeStr = fmt.Sprintf("%.2f", sizeF) + sizeMap[i+1]
	}
	return sizeStr
}
