package chat

import (
	"github.com/gofiber/websocket/v2"
)

type ChatClient struct {
	Conn  *websocket.Conn
	Email string
}
