package main

import (
	"domain/handler"
	"domain/service"
)

func main() {

	services := service.NewService()
	handlers := handler.NewHandler(services)

	router := handlers.SetupRouter()

	router.Run()
}
