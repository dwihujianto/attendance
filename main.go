package main

import (
	//"fmt"
	"api/app"
	"api/config"
)

func main() {
	config := config.GetConfig()

	//fmt.Println(config)
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
