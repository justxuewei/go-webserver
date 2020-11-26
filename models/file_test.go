package models

import (
	"github.com/xavier-niu/webserver/pkg/conf"
	"testing"
)

func TestListDir(t *testing.T) {
	items, _ := ListDir(conf.PublicDir)
	t.Log(items)
}
