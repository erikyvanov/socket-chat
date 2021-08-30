package controllers

import (
	"github.com/erikyvanov/chat-fh/jwt"
	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/services"
	"github.com/gofiber/fiber/v2"
)

func NewUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	userService := services.GetUserService()
	jwt, err := userService.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	return c.JSON(models.Response{Ok: true, Data: fiber.Map{
		"token": jwt,
		"user":  user,
	}})
}

func Login(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	userService := services.GetUserService()

	token, err := userService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	return c.JSON(models.Response{Ok: true, Data: fiber.Map{
		"token": token,
		"user":  user,
	}})
}

func Renew(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	token, err := jwt.GenerateJWT(*user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	userService := services.GetUserService()
	user, err = userService.GetUser(user.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	return c.JSON(models.Response{Ok: true, Data: fiber.Map{
		"token": token,
		"user":  user,
	}})
}
