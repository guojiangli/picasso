package main

import (
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/server/kgin"
	midd "github.com/guojiangli/picasso/pkg/server/kgin/middleware"
)

func main() {
	if err := Init(LoadWithFlags()); err != nil {
		panic(err)
	}
	defer func() {
		if msg := recover(); msg != nil {
			klog.Error("main-recover", msg)
		}
	}()

	var options, err = kgin.ConfigOption("picasso.server.gin.default")
	if err != nil {
		panic(err)
	}
	options.Logger = klog.DefaultLogger()
	s, err := kgin.NewServer(options)
	if err != nil {

		panic(err)
	}
	//新增系统recover
	s.Use(midd.Recovery(klog.DefaultLogger()))

	AddServer(s)
	//producer.InitMQProducer()
	//运行
	Run()

}
