package routes

import (
	"github.com/erikyvanov/chat-fh/controllers"
	"github.com/erikyvanov/chat-fh/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UsersGroup(app *fiber.App) {
	chatSocket := app.Group("/users", middlewares.ValidateUserJWT)
	chatSocket.Get("/all", controllers.GetAllUsers)

}
