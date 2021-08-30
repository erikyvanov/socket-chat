package middlewares

import (
	"errors"

	"github.com/erikyvanov/chat-fh/jwt"
	"github.com/erikyvanov/chat-fh/models"
	"github.com/gofiber/fiber/v2"
)

var ErrInvalidToken = errors.New("invalid token")

func ValidateUserJWT(c *fiber.Ctx) error {
	token := string(c.Request().Header.Peek("x-token"))
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Ok:   false,
			Data: "there is no token in the request",
		})
	}

	claims, err := jwt.TokenIsValid(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Ok:   false,
			Data: err.Error(),
		})
	}

	email, ok := claims["email"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Ok:   false,
			Data: ErrInvalidToken.Error(),
		})
	}

	name, ok := claims["name"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Ok:   false,
			Data: ErrInvalidToken.Error(),
		})
	}

	user := &models.User{
		Name:  name.(string),
		Email: email.(string),
	}

	c.Locals("user", user)
	return c.Next()
}
