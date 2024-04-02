package http

import (
	"net/http"
	"wifi_kost_be/helper"
	"wifi_kost_be/modules/usersubscription/service"

	"github.com/gofiber/fiber/v2"
)

type UserSubscriptionHandler struct {
	service service.UserSubscriptionService
}

func NewUserSubscriptionHandler(service service.UserSubscriptionService) *UserSubscriptionHandler {
	return &UserSubscriptionHandler{service: service}
}

func (h *UserSubscriptionHandler) GetUserSubscription(c *fiber.Ctx) error {
	type Request struct {
		UserID       int `json:"user_id"`
		GuestHouseID int `json:"guest_house_id"`
	}

	// Parse the JSON request
	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid JSON request", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Get the user_id from the request
	user_id := req.UserID
	guest_house_id := req.GuestHouseID
	guestHouses, err := h.service.GetUserSubscription(c.Context(), user_id, guest_house_id)

	if err != nil {
		// Handle the error
		response := helper.APIResponse("Failed to retrieve subscription detail", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if guestHouses == nil {
		// Handle case where no subscription are found
		response := helper.APIResponse("No package found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Return the subscription
	response := helper.APIResponse("UserSubscription Detail retrieved successfully", http.StatusOK, "OK", guestHouses)
	return c.Status(http.StatusOK).JSON(response)
}
