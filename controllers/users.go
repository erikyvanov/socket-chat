package controllers

import (
	"strconv"

	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: "invalid page"})
	}

	userService := services.GetUserService()
	users, err := userService.GetAllUsersExceptUserRequest(user.Email, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Ok: false, Errors: err.Error()})
	}

	return c.JSON(models.Response{Ok: true, Data: users})
}
