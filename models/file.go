package models

import (
	"os"
)

const (
	FileType = iota
	DirectoryType
)

type FileItem struct {
	ItemType int
	Name     string
}

func ListDir(path string) ([]FileItem, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return nil, err
	}

	items := make([]FileItem, 0)

	for _, fileInfo := range files {
		name := fileInfo.Name()
		itemType := FileType
		if fileInfo.IsDir() {
			itemType = DirectoryType
		}
		items = append(items, FileItem{ItemType: itemType, Name: name})
	}
	return items, nil
}
