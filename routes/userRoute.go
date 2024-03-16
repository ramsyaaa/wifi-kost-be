package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"wifi_kost_be/modules/user/http"
	"wifi_kost_be/modules/user/repository"
	"wifi_kost_be/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func UserRouter(app *fiber.App, db *gorm.DB) {
	jwtSecret := os.Getenv("JWT_SECRET")
	ctx := context.Background()
	var (
		// redis configs
		redisHost     = os.Getenv("REDIS_HOST")
		redisPort     = os.Getenv("REDIS_PORT")
		redisPassword = os.Getenv("REDIS_PASSWORD")
	)
	rds := setupRedis(ctx, redisHost, redisPort, redisPassword)
	fmt.Println(rds)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	tokenService := service.NewTokenService(jwtSecret)
	userHandler := http.NewUserHandler(userService, tokenService, rds)

	http.UserRoutes(app, userHandler)

}

func setupRedis(ctx context.Context, host string, port string, password string) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	_, err := rds.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to Redis..")
	return rds
}
