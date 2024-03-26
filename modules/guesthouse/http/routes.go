package http

import (
	"github.com/gofiber/fiber/v2"
)

func GuestHouseRoutes(app *fiber.App, handler *GuestHouseHandler) {

	app.Get("/guest-house/list", handler.GetGuestHouse)
	app.Post("/guest-house/detail", handler.GetGuestHouseDetail)
}
