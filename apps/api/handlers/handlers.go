package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type hHandlers struct {
}

func NewHandlers() Handlers {
	return &hHandlers{}
}

func (h *hHandlers) Hello(c *fiber.Ctx) error {
	return c.Status(200).SendString("hello truck unii")
}
