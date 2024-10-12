package main

import (
	"example/controller"
	"example/repository"
	"example/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repo := repository.New()
	service := service.New(repo)
	controller := controller.New(service)

	app := fiber.New()

	app.Post("/login", controller.Login)

	app.Listen(":8080")
}
