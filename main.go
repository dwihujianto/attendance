package main

import (
	"fmt"
	"github.com/dwihujianto/attendance/app"
	"github.com/dwihujianto/attendance/config"
)

func main() {
	config := config.GetConfig()

	fmt.Println(config.DB.Dialect)
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
