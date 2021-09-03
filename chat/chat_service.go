package chat

import (
	"sync"

	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/repositories"
	"github.com/erikyvanov/chat-fh/services"
	"github.com/gofiber/websocket/v2"
)

var onceChatService sync.Once
var chatServiceSingleton *ChatService

type ChatService struct {
	users              map[string]*websocket.Conn
	UpcomingChatClient chan ChatClient
	DeleteChatClient   chan ChatClient
	UpcomingMessage    chan models.ChatMessage
}

func (cs *ChatService) Run() {
	messagesRepository := repositories.GetMessageRepository()
	userService := services.GetUserService()

	for {
		select {
		case upcomingChatClient := <-cs.UpcomingChatClient:
			cs.users[upcomingChatClient.Email] = upcomingChatClient.Conn

		case upcomingMessage := <-cs.UpcomingMessage:
			newID, err := messagesRepository.SaveMessage(upcomingMessage)
			if err != nil {
				return
			}

			upcomingMessage.ID = newID

			if c, ok := cs.users[upcomingMessage.ReciverEmail]; ok {
				err := c.WriteJSON(upcomingMessage)
				if err != nil {
					delete(cs.users, upcomingMessage.ReciverEmail)
				}
			}

		case deleteChatClient := <-cs.DeleteChatClient:
			delete(cs.users, deleteChatClient.Email)
			go userService.SetConnectionStatus(deleteChatClient.Email, true)
		}
	}
}

func GetChatService() *ChatService {
	onceChatService.Do(func() {
		chatServiceSingleton = &ChatService{
			users:              make(map[string]*websocket.Conn),
			UpcomingChatClient: make(chan ChatClient),
			DeleteChatClient:   make(chan ChatClient),
			UpcomingMessage:    make(chan models.ChatMessage),
		}
	})

	return chatServiceSingleton
}
