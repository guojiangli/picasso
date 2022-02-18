package main

import (
	"fmt"
	"github.com/guojiangli/picasso/core"
	"github.com/guojiangli/picasso/pkg/klog"
	"github.com/guojiangli/picasso/pkg/server/kgin"
	midd "github.com/guojiangli/picasso/pkg/server/kgin/middleware"
)

func main() {
	if err := core.Init(core.LoadWithFlags()); err != nil {
		fmt.Println("ddd")
		panic(err)
	}
	defer func() {
		if msg := recover(); msg != nil {
			klog.Error("main-recover：", msg)
		}
	}()

	var options, err = kgin.ConfigOption("picasso.server.gin.default")
	if err != nil {
		fmt.Println("对对滴")
		panic(err)
	}
	options.Logger = klog.DefaultLogger()
	s, err := kgin.NewServer(options)
	if err != nil {
		fmt.Println("对对滴  ddd ")
		panic(err)
	}
	//新增系统recover
	s.Use(midd.Recovery(klog.DefaultLogger()))
	fmt.Println("对对滴  ddd dddddff")
	core.AddServer(s)
	fmt.Println("对对滴  ddd ")
	//producer.InitMQProducer()
	//运行
	core.Run()
	fmt.Println("对对滴  ddd ")

}
