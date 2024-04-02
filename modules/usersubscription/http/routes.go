package http

import (
	"github.com/gofiber/fiber/v2"
)

func UserSubscriptionRoutes(app *fiber.App, handler *UserSubscriptionHandler) {

	app.Post("/user/subscription", handler.GetUserSubscription)
}
