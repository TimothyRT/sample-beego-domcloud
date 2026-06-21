package main

import (
	"os"
	_ "sample-beego-domcloud/routers"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if portStr := os.Getenv("PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			beego.BConfig.Listen.HTTPPort = port
		}
	}
	beego.Run()
}
