package conf

import (
	"fmt"
	"github.com/xavier-niu/webserver/pkg/util"
)

var Host = ""
var Port = 10000
var Addr = fmt.Sprintf("%s:%d", Host, Port)

// built-in directories
var HomeDir = util.GetHomeDirectory()
var PublicDir = HomeDir + "/Public"
