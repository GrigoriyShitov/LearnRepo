package main

import (
	"awesome-backend/internal/app"
	"awesome-backend/internal/config"
	"awesome-backend/internal/container"

	"git.itmo.su/go-modules/log/v4"
)

const (
	appName = "awesome-backend"
)

// @title           awesome-backend API
// @version         0.0.1
// @description     REST API
// @BasePath        /awesome-backend/api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitApp()
	log.InitLogging(appName, config.App.Env)
	//xerr.StartErrorCodesFrom(config.App.StartErrorCodesFrom)

	c := container.New(appName)
	a := app.New(c.Server(), c.Starters(), c.Stoppers())
	err := a.Init()
	if err != nil {
		log.Error(err)
	}
}
