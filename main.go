package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"kingProject/app/api/coinMonitor/alcor"
	"kingProject/app/api/coinMonitor/coingecko"
	_ "kingProject/router"
)

func main() {
	gcron.SetLogLevel(glog.LEVEL_ALL)
	alcor.MonitorPrice()
	coingecko.MonitorPrice()
	g.Server().Run()
}
