package http

import (
	"github.com/gofiber/fiber/v2"
)

func PackageRoutes(app *fiber.App, handler *PackageHandler) {

	app.Post("/package/list", handler.GetPackage)
	app.Post("/package/detail", handler.GetPackageDetail)
}
