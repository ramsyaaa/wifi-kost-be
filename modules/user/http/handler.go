package http

import (
	"net/http"
	"time"
	"wifi_kost_be/helper"
	"wifi_kost_be/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service      service.UserService
	tokenService service.TokenService
	redisClient  *redis.Client
}

func NewUserHandler(service service.UserService, tokenService service.TokenService, redisClient *redis.Client) *UserHandler {
	return &UserHandler{service: service, tokenService: tokenService, redisClient: redisClient}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Msisdn       string `json:"msisdn"`
		Password     string `json:"password"`
		GuestHouseID int    `json:"guest_house_id"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	guest_house_id := req.GuestHouseID

	// Authenticate the user
	user, err := h.service.FindByMsisdn(c.Context(), req.Msisdn)
	if err != nil {
		return err
	}

	if user == nil {
		response := helper.APIResponse("User Not Found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Compare the hashed password from the database with the password provided in the request
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response := helper.APIResponse("Wrong Password", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Passwords match, create a JWT token for the user
	tokenString, err := h.tokenService.GenerateToken(user, guest_house_id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "OK", tokenString)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	response := helper.APIResponse("Logout Success", http.StatusOK, "OK", nil)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Msisdn       string `json:"msisdn"`
		GuestHouseID int    `json:"guest_house_id"`
	}
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	msisdn := req.Msisdn
	guest_house_id := req.GuestHouseID
	user, err := h.service.FindByMsisdn(c.Context(), msisdn)
	if err != nil {
		return err
	}

	// Check if the user's msisdn is already registered

	if user != nil {
		response := helper.APIResponse("User Already Exist", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	tokenString, err := h.tokenService.GenerateGuestToken(msisdn, guest_house_id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if err := h.redisClient.Set(c.Context(), msisdn, tokenString, time.Hour); err == nil {
		response := helper.APIResponse("Failed to register user", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse("User registered successfully", http.StatusOK, "OK", tokenString)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) CheckGuest(c *fiber.Ctx) error {
	// Define a struct to represent the JSON request
	type Request struct {
		Key string `json:"key"`
	}

	// Parse the JSON request
	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid JSON request", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Get the key from the request
	key := req.Key

	// Get data from Redis
	value, err := h.redisClient.Get(c.Context(), key).Result()
	if err != nil {
		if err == redis.Nil {
			// Key does not exist in Redis
			response := helper.APIResponse("Key does not exist in Redis", http.StatusNotFound, "Error", nil)
			return c.Status(http.StatusNotFound).JSON(response)
		}
		// Other error occurred
		response := helper.APIResponse("Failed to get data from Redis", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Data exists in Redis
	response := helper.APIResponse("Data found in Redis", http.StatusOK, "OK", value)
	return c.Status(http.StatusOK).JSON(response)
}
