package conf

import (
	"os"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

func setPort() {
	if portStr := os.Getenv("PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			beego.BConfig.Listen.HTTPPort = port
		}
	}
}

func init() {
	setPort()
}
