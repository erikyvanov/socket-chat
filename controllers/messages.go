package controllers

import (
	"strconv"

	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetChatMessages(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	emailUserChat := c.Params("user_chat")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{Ok: false, Errors: "invalid page"})
	}

	messageRepository := repositories.GetMessageRepository()
	messages, err := messageRepository.GetChatMessages(user.Email, emailUserChat, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{Ok: false, Errors: "error in database"})
	}

	return c.JSON(models.Response{Ok: true, Data: messages})
}
