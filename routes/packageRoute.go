package routes

import (
	"wifi_kost_be/modules/package/http"
	"wifi_kost_be/modules/package/repository"
	"wifi_kost_be/modules/package/service"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func PackageRouter(app *fiber.App, db *gorm.DB) {
	packageRepo := repository.NewPackageRepository(db)
	packageService := service.NewPackageService(packageRepo)
	packageHandler := http.NewPackageHandler(packageService)

	http.PackageRoutes(app, packageHandler)

}
