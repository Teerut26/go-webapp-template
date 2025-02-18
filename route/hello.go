package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.SendString(uuid.New().String())
}
