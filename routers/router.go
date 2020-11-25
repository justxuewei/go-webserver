package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xavier-niu/webserver/controllers"
	"github.com/xavier-niu/webserver/pkg/util"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	InitNetDiskRouter(r)
	util.Log().Info("NetDisk routers are initialized.")
	return r
}

func InitNetDiskRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("public", controllers.Public)
}
