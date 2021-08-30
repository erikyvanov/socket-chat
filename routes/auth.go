package routes

import (
	"github.com/erikyvanov/chat-fh/controllers"
	"github.com/erikyvanov/chat-fh/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthGroup(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/new", middlewares.ValidateUserFields, controllers.NewUser)
	auth.Post("/login", middlewares.ValidateLoginFields, controllers.Login)
	auth.Get("/renew", middlewares.ValidateUserJWT, controllers.Renew)
}
