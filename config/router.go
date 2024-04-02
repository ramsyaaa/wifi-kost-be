package config

import (
	"log"
	"os"
	"wifi_kost_be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) {

	app := fiber.New()
	// Use the cors middleware to allow all origins and methods
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Create a new Fiber app for the "api/v1" prefix group
	api := fiber.New()

	routes.UserRouter(api, db)
	routes.GuestHouseRouter(api, db)
	routes.PackageRouter(api, db)
	routes.TransactionRouter(api, db)
	routes.UserSubscriptionRouter(api, db)

	// Mount the "api/v1" group under the main app
	app.Mount("/api/v1", api)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
