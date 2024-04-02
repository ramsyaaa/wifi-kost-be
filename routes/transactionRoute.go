package routes

import (
	"wifi_kost_be/modules/transaction/http"
	"wifi_kost_be/modules/transaction/repository"
	"wifi_kost_be/modules/transaction/service"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func TransactionRouter(app *fiber.App, db *gorm.DB) {
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := http.NewTransactionHandler(transactionService)

	http.TransactionRoutes(app, transactionHandler)

}
