package netdisk

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xavier-niu/webserver/models"
	"github.com/xavier-niu/webserver/pkg/util"
	"net/url"
	"strings"
)

type ListService struct{}

type DownloadService struct{}

type FileItem struct {
	ItemType int
	Name     string
	Url      string
}

func (l *ListService) List(c *gin.Context, path string) (int, string, *gin.H) {
	items, err := models.ListDir(path)
	if err != nil {
		util.Log().Error("Listing \"%s\" is failed, %s.", path, err)
		return util.ErrorPage("List Error", "The directory could not be accessed.")
	}

	dirs := make([]FileItem, 0)
	for _, item := range items {
		name := item.Name
		u := c.Request.RequestURI + name
		itemType := item.ItemType
		dirs = append(dirs, FileItem{ItemType: itemType, Name: name, Url: u})
	}

	dirName, _ := url.QueryUnescape(util.TrimFirstRune(c.Request.RequestURI))
	return 200, "netdisk/list.html", &gin.H{
		"title":   fmt.Sprintf("%s - NetDisk", dirName),
		"dirName": dirName,
		"dirs":    dirs,
	}
}

func (dl *DownloadService) Download(c *gin.Context, path string) {
	pathArr := strings.Split(path, "/")
	filename := pathArr[len(pathArr)-1]
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(path)
}
