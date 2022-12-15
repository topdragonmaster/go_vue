package main

import (
	"sniper/door/config"

	"sniper/door/app"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
