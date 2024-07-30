package main

import (
	"github.com/TimurZheksimbaev/Golang-webchat/config"
	"github.com/TimurZheksimbaev/Golang-webchat/database"
	"github.com/TimurZheksimbaev/Golang-webchat/server/routers"
	"github.com/TimurZheksimbaev/Golang-webchat/server/user"
	"github.com/TimurZheksimbaev/Golang-webchat/server/ws"
	"github.com/TimurZheksimbaev/Golang-webchat/utils"
)

func main() {
	config, err := config.LoadEnv()
	utils.LogExit(err)

	database, dbErr := database.NewDatabase(config)
	utils.LogExit(dbErr)

	userRepo := user.NewRepository(database.GetDB())
	userService := user.NewService(userRepo, config)
	userHandler := user.NewHandler(userService)


	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()


	r := routers.InitRouter(config,userHandler, wsHandler)
	r.Run(config.ServerHost + ":" + config.ServerPort)

}