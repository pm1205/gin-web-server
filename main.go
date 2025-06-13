package main

import (
	"gin-web-server/config"
	"gin-web-server/route"
)

func main() {
	config.InitDB()
	r := route.SetupRouter()
	r.Run(":8080")
}
