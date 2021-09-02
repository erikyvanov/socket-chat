package routes

import (
	"github.com/erikyvanov/chat-fh/controllers"
	"github.com/erikyvanov/chat-fh/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MessagesGroup(app *fiber.App) {
	messagesGroup := app.Group("/messages")
	messagesGroup.Use(middlewares.ValidateUserJWT)
	messagesGroup.Get("/:user_chat", controllers.GetChatMessages)
}
