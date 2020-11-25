package main

import (
	"github.com/xavier-niu/webserver/pkg/conf"
	"github.com/xavier-niu/webserver/pkg/util"
	"github.com/xavier-niu/webserver/routers"
)

func main() {
	app := routers.InitRouter()

	ipAddr, err := util.GetLocalIpAddr()
	if err != nil {
		util.Log().Error("Getting IP address is failed, %s", err)
	}

	util.Log().Info("Go-webserver can be accessed by %s:%d.", ipAddr, conf.Port)
	if err := app.Run(conf.Addr); err != nil {
		util.Log().Error("go-webserver is failed to run, %s", err)
		return
	}
}
