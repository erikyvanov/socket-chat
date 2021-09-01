package models

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

var onceChatService sync.Once
var chatServiceSingleton *ChatService

type ChatService struct {
	users              map[string]*websocket.Conn
	UpcomingChatClient chan ChatClient
	DeleteChatClient   chan ChatClient
	UpcomingMessage    chan ChatMessage
}

func (cs *ChatService) Run() {
	for {
		select {
		case upcomingChatClient := <-cs.UpcomingChatClient:
			cs.users[upcomingChatClient.Email] = upcomingChatClient.Conn

		case upcomingMessage := <-cs.UpcomingMessage:
			if c, ok := cs.users[upcomingMessage.EmailReciver]; ok {
				err := c.WriteJSON(upcomingMessage)
				if err != nil {

					delete(cs.users, upcomingMessage.EmailReciver)
				}
			}

		case deleteChatClient := <-cs.DeleteChatClient:
			delete(cs.users, deleteChatClient.Email)
		}
	}
}

func GetChatService() *ChatService {
	onceChatService.Do(func() {
		chatServiceSingleton = &ChatService{
			users:              make(map[string]*websocket.Conn),
			UpcomingChatClient: make(chan ChatClient),
			DeleteChatClient:   make(chan ChatClient),
			UpcomingMessage:    make(chan ChatMessage),
		}
	})

	return chatServiceSingleton
}
