package middlewares

import (
	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/utils"
	"github.com/gofiber/fiber/v2"
)

func ValidateUserFields(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	errors := utils.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: errors})
	}

	c.Locals("user", user)
	return c.Next()
}

func ValidateLoginFields(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	errors := utils.ValidateLogin(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: errors})
	}

	c.Locals("user", user)
	return c.Next()
}
