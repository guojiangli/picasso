package dialector

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dialector struct {
	gorm.Dialector
}

func New(opts ...*Option) *Dialector {
	DialectorOpts := defaultOption().MergeOption(opts...)
	Dialector := Dialector{}
	Dialector.Dialector = mysql.Open(DialectorOpts.DSN)
	return &Dialector
}
