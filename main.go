package main

import (
	"github.com/erikyvanov/chat-fh/database"
	"github.com/erikyvanov/chat-fh/routes"
	"github.com/erikyvanov/chat-fh/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnvFile()
	if database.IsConnectedToMongoDB() {
		defer database.CloseConnection()
		app := fiber.New()
		routes.AuthGroup(app)

		app.Listen(":3000")
	}
}
