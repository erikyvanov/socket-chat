package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func IsWebSocketUpgrade(c *fiber.Ctx) error {

	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)

		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
