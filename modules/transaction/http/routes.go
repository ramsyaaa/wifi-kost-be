package http

import (
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App, handler *TransactionHandler) {

	app.Post("/transaction/create", handler.CreateTransaction)
}
