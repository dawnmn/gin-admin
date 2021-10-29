package app

import (
	"gin-admin/config"
)

func init() {
	config.InitConfig(&Config)
	InitAppLog()
	InitRouter()
}
