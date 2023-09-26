package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	Hello(*fiber.Ctx) error
}

type HandlerRecycle interface {
}

type HandlerTruck interface {
	TruckProfileGet(*fiber.Ctx) error
}
