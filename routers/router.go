package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xavier-niu/webserver/controllers"
	"github.com/xavier-niu/webserver/pkg/util"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// static resources
	r.LoadHTMLGlob("templates/**/*")

	InitNetDiskRouter(r)
	util.Log().Info("NetDisk routers are initialized.")
	return r
}

func InitNetDiskRouter(r *gin.Engine) {
	// router
	{
		// public directory
		r.GET("/public/*subpath", controllers.Public)
	}
}
