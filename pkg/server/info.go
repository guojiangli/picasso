package server

import (
	"fmt"
)

type ServiceInfo struct {
	Name string
	Host string
	Port int
}

func (info *ServiceInfo) Address() string {
	return fmt.Sprintf("%s:%d", info.Host, info.Port)
}
