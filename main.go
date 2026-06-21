package main

import (
	_ "sample-beego-domcloud/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	fmt.Println("PORT =", os.Getenv("PORT"))
	beego.Run()
}

