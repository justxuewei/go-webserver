package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xavier-niu/webserver/pkg/conf"
	"github.com/xavier-niu/webserver/pkg/util"
	"github.com/xavier-niu/webserver/service/netdisk"
	"os"
)

func Public(c *gin.Context) {
	var listService netdisk.ListService
	var downloadService netdisk.DownloadService

	path := conf.PublicDir
	subPath := c.Param("subpath")
	if subPath != "/" {
		path += subPath
	}
	f, err := os.Stat(path)
	if err != nil {
		util.Log().Debug(err.Error())
		code, tmpl, h := util.ErrorPage("File Error", "Fail to open the path")
		c.HTML(code, tmpl, *h)
		return
	}
	if f.IsDir() {
		code, tmpl, h := listService.List(c, path)
		c.HTML(code, tmpl, *h)
	} else {
		downloadService.Download(c, path)
	}
}
