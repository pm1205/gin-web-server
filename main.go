package main

import (
	"server/config"
	"server/model"
	"server/route"
)

func main() {
	config.InitDB()

	// 自动迁移模型
	config.DB.AutoMigrate(&model.User{})

	r := route.SetupRouter()
	r.Run(":8080")
}
