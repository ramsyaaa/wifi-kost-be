package http

import (
	"wifi_kost_be/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *UserHandler) {

	app.Post("/login", handler.Login)
	app.Post("/register", handler.Register)
	app.Get("/check-guest", handler.CheckGuest)
	app.Post("/logout", middleware.AuthMiddleware(), handler.Logout)
}
