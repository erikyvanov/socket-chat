package models

import (
	"github.com/gofiber/websocket/v2"
)

type ChatClient struct {
	Conn  *websocket.Conn
	Email string
}

// func (cc *ChatClient) Run() {
// 	for {
// 		select {
// 		case msg := <-cc.UpcomingMsg:
// 			cc.Conn.WriteMessage(websocket.TextMessage, []byte(msg.Message))

// 		case <-cc.close:
// 			cc.Conn.Close()
// 		}
// 	}
// }
