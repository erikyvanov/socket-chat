package controllers

import (
	"github.com/erikyvanov/chat-fh/chat"
	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/services"
	"github.com/gofiber/websocket/v2"
)

func IntoToChat(c *websocket.Conn) {
	var message models.ChatMessage

	user := c.Locals("user").(*models.User)
	userService := services.GetUserService()

	err := userService.SetConnectionStatus(user.Email, true)
	if err != nil {
		c.Close()
		return
	}

	chatService := chat.GetChatService()
	chatClient := chat.ChatClient{Conn: c, Email: user.Email}

	chatService.UpcomingChatClient <- chatClient

	for {

		err := c.ReadJSON(&message)
		if err != nil {
			chatService.DeleteChatClient <- chatClient
			break
		}

		chatService.UpcomingMessage <- message

	}
}
