package routes

import (
	"github.com/erikyvanov/chat-fh/controllers"
	"github.com/erikyvanov/chat-fh/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ChatSocketGroup(app *fiber.App) {
	chatSocket := app.Group("/ws")
	chatSocket.Use("/chat", middlewares.IsWebSocketUpgrade, middlewares.ValidateUserJWT)
	chatSocket.Get("/chat", websocket.New(controllers.IntoToChat))

}
