package app

import (
	"github.com/dwihujianto/attendance/app/handler/employee"
	"github.com/dwihujianto/attendance/app/model"
	"github.com/dwihujianto/attendance/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize(config *config.Config) {
	model.DBInit(config)
	r := gin.Default()
	a.Router = r
	a.setRouters()
}

func (a *App) setRouters() {
	api := a.Router.Group("api")

	api.GET("/employees", employee.GetAll)
	api.GET("/employees/:id", employee.GetById)
}

func (a *App) Run(host string) {
	a.Router.Run(host)
}
