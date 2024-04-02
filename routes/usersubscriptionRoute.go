package routes

import (
	"wifi_kost_be/modules/usersubscription/http"
	"wifi_kost_be/modules/usersubscription/repository"
	"wifi_kost_be/modules/usersubscription/service"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func UserSubscriptionRouter(app *fiber.App, db *gorm.DB) {
	usersubscriptionRepo := repository.NewUserSubscriptionRepository(db)
	usersubscriptionService := service.NewUserSubscriptionService(usersubscriptionRepo)
	usersubscriptionHandler := http.NewUserSubscriptionHandler(usersubscriptionService)

	http.UserSubscriptionRoutes(app, usersubscriptionHandler)

}
