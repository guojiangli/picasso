package kgorm

import (
	"sync"
)

var _instances = sync.Map{}

// Get ...
func Get(name string) (*DB, bool) {
	if ins, ok := _instances.Load(name); ok {
		return ins.(*DB), true
	}
	return nil, false
}
