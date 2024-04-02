package http

import (
	"net/http"
	"strings"
	"time"
	"wifi_kost_be/helper"
	"wifi_kost_be/models"
	"wifi_kost_be/modules/transaction/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	// Parse request body
	var req models.Transaction
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	var request models.UserRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	// Check if msisdn is provided, if not, it means the transaction is for a registered user
	if request.Msisdn == "" && req.UserID == 0 {
		response := helper.APIResponse("Invalid request: msisdn or user_id is required", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// If msisdn is provided, format it to 628xx if it's in 08xx format
	if request.Msisdn != "" {
		request.Msisdn = formatMsisdn(request.Msisdn)
	}

	// Check if the user already exists
	var user *models.User
	var err error // Declare err outside to avoid redeclaration
	if req.UserID != 0 {
		user, err = h.service.FindById(c.Context(), req.UserID) // Assign value to existing user variable
		if err != nil {
			response := helper.APIResponse("Failed to get user", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
		if user == nil {
			response := helper.APIResponse("User not found", http.StatusBadRequest, "Error", nil)
			return c.Status(http.StatusBadRequest).JSON(response)
		}
		transaction := models.Transaction{
			UserID:            req.UserID,
			PackageID:         req.PackageID,
			GuestHouseID:      req.GuestHouseID,
			Amount:            req.Amount,
			TotalAmount:       req.TotalAmount,
			IsUsingCoupon:     req.IsUsingCoupon,
			PaymentMethod:     req.PaymentMethod,
			TransactionStatus: "settled",
			TransactionDate:   time.Now(),
		}

		subscription := models.UserSubscription{
			UserID:       req.UserID,
			PackageID:    req.PackageID,
			GuestHouseID: req.GuestHouseID,
		}
		// Create transaction using the service
		if err := h.service.CreateTransaction(c.Context(), &transaction); err != nil {
			response := helper.APIResponse("Failed to create transaction", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}

		if err := h.service.CreateSubscription(c.Context(), &subscription); err != nil {
			response := helper.APIResponse("Failed to create transaction", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
	} else {
		// Create the user if it doesn't exist
		user, err = h.service.CreateUser(c.Context(), request.Msisdn) // Assign value to existing user variable
		if err != nil {
			response := helper.APIResponse("Failed to create user", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
		req.UserID = user.ID
		// Fetch the guest house name and package details
		guestHouse, err := h.service.GetGuestHouseDetail(c.Context(), req.GuestHouseID)
		if err != nil {
			response := helper.APIResponse("Failed to get guest house detail", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}

		packageDetail, err := h.service.GetPackageDetail(c.Context(), req.PackageID)
		if err != nil {
			response := helper.APIResponse("Failed to get package detail", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}

		// Send welcome SMS to the newly created user
		if err := helper.SendSMS(user.Msisdn, guestHouse.GuestHouseName, packageDetail.ExpiryDay); err != nil {
			response := helper.APIResponse("Failed to send welcome SMS", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
		transaction := models.Transaction{
			UserID:            user.ID,
			PackageID:         req.PackageID,
			GuestHouseID:      req.GuestHouseID,
			Amount:            req.Amount,
			TotalAmount:       req.TotalAmount,
			IsUsingCoupon:     req.IsUsingCoupon,
			PaymentMethod:     req.PaymentMethod,
			TransactionStatus: "settled",
			TransactionDate:   time.Now(),
		}
		subscription := models.UserSubscription{
			UserID:       user.ID,
			PackageID:    req.PackageID,
			GuestHouseID: req.GuestHouseID,
		}
		// Create transaction using the service
		if err := h.service.CreateTransaction(c.Context(), &transaction); err != nil {
			response := helper.APIResponse("Failed to create transaction", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}

		if err := h.service.CreateSubscription(c.Context(), &subscription); err != nil {
			response := helper.APIResponse("Failed to create transaction", http.StatusInternalServerError, "Error", nil)
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
	}

	// Respond with success message
	response := helper.APIResponse("Transaction created successfully", http.StatusOK, "OK", nil)
	return c.Status(http.StatusOK).JSON(response)
}

func formatMsisdn(msisdn string) string {
	if strings.HasPrefix(msisdn, "08") {
		return "62" + msisdn[1:]
	}
	return msisdn
}
