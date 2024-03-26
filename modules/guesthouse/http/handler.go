package http

import (
	"net/http"
	"wifi_kost_be/helper"
	"wifi_kost_be/modules/guesthouse/service"

	"github.com/gofiber/fiber/v2"
)

type GuestHouseHandler struct {
	service service.GuestHouseService
}

func NewGuestHouseHandler(service service.GuestHouseService) *GuestHouseHandler {
	return &GuestHouseHandler{service: service}
}

func (h *GuestHouseHandler) GetGuestHouse(c *fiber.Ctx) error {
	guestHouses, err := h.service.GetGuestHouse(c.Context())
	if err != nil {
		// Handle the error
		response := helper.APIResponse("Failed to retrieve guest houses", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if guestHouses == nil {
		// Handle case where no guest houses are found
		response := helper.APIResponse("No guest houses found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Return the guest houses
	response := helper.APIResponse("Guest houses retrieved successfully", http.StatusOK, "OK", guestHouses)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *GuestHouseHandler) GetGuestHouseDetail(c *fiber.Ctx) error {
	type Request struct {
		GuestHouseID int `json:"guest_house_id"`
	}

	// Parse the JSON request
	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid JSON request", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	guest_house_id := req.GuestHouseID
	guestHouse, err := h.service.GetGuestHouseDetail(c.Context(), guest_house_id)
	if err != nil {
		// Handle the error
		response := helper.APIResponse("Failed to retrieve guest houses", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if guestHouse == nil {
		// Handle case where no guest houses are found
		response := helper.APIResponse("No guest houses found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Return the guest houses
	response := helper.APIResponse("Guest houses retrieved successfully", http.StatusOK, "OK", guestHouse)
	return c.Status(http.StatusOK).JSON(response)
}
