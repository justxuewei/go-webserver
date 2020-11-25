package netdisk

import (
	"github.com/gin-gonic/gin"
	"github.com/xavier-niu/webserver/pkg/conf"
	"github.com/xavier-niu/webserver/pkg/util"
	"os"
	"path/filepath"
)

const (
	FileType = iota
	DirectoryType
)

type ListService struct {}

type FileItem struct {
	itemType int
	path string
}

func (l *ListService) ListPublic(c *gin.Context) (string, *gin.H) {
	pubDirInfo, err := os.Stat(conf.PublicDir)
	if os.IsNotExist(err) {
		h := gin.H{
			"title": "No Authorization",
			"error": "Public folder is not existed.",
		}
		return "error.html", &h
	}
	if !pubDirInfo.IsDir() {
		h := gin.H{
			"title": "Type Error",
			"error": "Public should be a directory.",
		}
		return "error.html", &h
	}
	items := make([]*FileItem, 10)
	visitFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			util.Log().Error("Visiting \"%s\" gets an error, %s", path, err)
			return err
		}
		itemType := FileType
		if info.IsDir() {
			itemType = DirectoryType
		}
		item := FileItem{itemType: itemType, path: path}
		items = append(items, &item)
		return nil
	}
	err := filepath.Walk(conf.PublicDir, visitFunc)
	if err != nil {

	}
}
