package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	msisdn := req.Msisdn

	// Validate the format of the msisdn
	if !isValidMsisdnFormat(msisdn) {
		response := helper.APIResponse("Invalid msisdn format", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Format the msisdn to 628xx if it's in 08xx format
	msisdn = formatMsisdn(msisdn)
	// Authenticate the user
	user, err := h.service.FindByMsisdn(c.Context(), msisdn)
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
	guestHouseID := req.GuestHouseID

	// Validate the format of the msisdn
	if !isValidMsisdnFormat(msisdn) {
		response := helper.APIResponse("Invalid msisdn format", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Format the msisdn to 628xx if it's in 08xx format
	msisdn = formatMsisdn(msisdn)

	// Check if the user's msisdn is already registered in the database
	user, err := h.service.FindByMsisdn(c.Context(), msisdn)
	if err != nil {
		response := helper.APIResponse("Failed to get user", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	if user != nil {
		response := helper.APIResponse("User already exists", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Check if the msisdn is already registered with the external API
	externalCheckURL := fmt.Sprintf("https://g2-radius-api.aneta.my.id/user-username/%s", msisdn)
	checkResp, err := http.Get(externalCheckURL)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	defer checkResp.Body.Close()

	// Read the response from the external API
	checkBody, err := ioutil.ReadAll(checkResp.Body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Parse the response
	var existingUsers []interface{}
	if err := json.Unmarshal(checkBody, &existingUsers); err != nil {
		return fiber.ErrInternalServerError
	}

	// If the response is not empty, the user already exists
	if len(existingUsers) > 0 {
		response := helper.APIResponse("User already exists in external service", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Create the user in the database
	user, err = h.service.CreateUser(c.Context(), msisdn)
	if err != nil {
		response := helper.APIResponse("Failed to create user in database", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	// Prepare the request payload for the external API
	externalRequest := map[string]string{
		"username":  msisdn,
		"value":     "password",
		"firstname": "-",
		"lastname":  "-",
		"groupname": "wifi-kost",
	}

	externalRequestBody, err := json.Marshal(externalRequest)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Call the external API to register the user
	externalAPIURL := "https://g2-radius-api.aneta.my.id/user-create"
	resp, err := http.Post(externalAPIURL, "application/json", bytes.NewBuffer(externalRequestBody))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	defer resp.Body.Close()

	// Check the response from the external API
	if resp.StatusCode != http.StatusOK {
		response := helper.APIResponse("Failed to register user with external service", resp.StatusCode, "Error", nil)
		return c.Status(resp.StatusCode).JSON(response)
	}

	// Generate the token
	tokenString, err := h.tokenService.GenerateGuestToken(msisdn, guestHouseID)
	if err != nil {
		return fiber.ErrInternalServerError
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

	// // Get data from Redis
	// value, err := h.redisClient.Get(c.Context(), key).Result()
	// if err != nil {
	// 	if err == redis.Nil {
	// 		// Key does not exist in Redis
	// 		response := helper.APIResponse("Key does not exist in Redis", http.StatusNotFound, "Error", nil)
	// 		return c.Status(http.StatusNotFound).JSON(response)
	// 	}
	// 	// Other error occurred
	// 	response := helper.APIResponse("Failed to get data from Redis", http.StatusInternalServerError, "Error", nil)
	// 	return c.Status(http.StatusInternalServerError).JSON(response)
	// }

	// Data exists in Redis
	response := helper.APIResponse("Data found in Redis", http.StatusOK, "OK", key)
	return c.Status(http.StatusOK).JSON(response)
}

func isValidMsisdnFormat(msisdn string) bool {
	return strings.HasPrefix(msisdn, "08") || strings.HasPrefix(msisdn, "628")
}

func formatMsisdn(msisdn string) string {
	if strings.HasPrefix(msisdn, "08") {
		return "62" + msisdn[1:]
	}
	return msisdn
}
