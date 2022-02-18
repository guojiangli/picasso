package rotate

import (
	"picasso/pkg/klog/rotate/lumberjack"
)

func NewRotate(opts ...*RotateOption) *lumberjack.Logger {
	RotateOpt := defaultOption().MergeOption(opts...)
	return &lumberjack.Logger{
		Filename:   RotateOpt.FileName,
		MaxSize:    RotateOpt.MaxSize,
		MaxAge:     RotateOpt.MaxAge,
		MaxBackups: RotateOpt.MaxBackups,
		LocalTime:  RotateOpt.LocalTime,
		Compress:   RotateOpt.Compress,
	}
}
