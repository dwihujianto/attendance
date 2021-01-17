package main

import (
	"github.com/dwihujianto/attendance/app"
	"github.com/dwihujianto/attendance/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
