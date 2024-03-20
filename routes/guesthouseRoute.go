package routes

import (
	"wifi_kost_be/modules/guesthouse/http"
	"wifi_kost_be/modules/guesthouse/repository"
	"wifi_kost_be/modules/guesthouse/service"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func GuestHouseRouter(app *fiber.App, db *gorm.DB) {
	guesthouseRepo := repository.NewGuestHouseRepository(db)
	guesthouseService := service.NewGuestHouseService(guesthouseRepo)
	guesthouseHandler := http.NewGuestHouseHandler(guesthouseService)

	http.GuestHouseRoutes(app, guesthouseHandler)

}
