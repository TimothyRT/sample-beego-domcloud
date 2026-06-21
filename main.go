package main

import (
	"fmt"
	"os"
	_ "sample-beego-domcloud/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if portStr := os.Getenv("PORT"); portStr != "" {
        port, err := strconv.Atoi(portStr)
        if err == nil {
            web.BConfig.Listen.HTTPPort = port
        }
    }
	fmt.Println("PORT =", os.Getenv("PORT"))
	beego.Run()
}

