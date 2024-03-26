package http

import (
	"net/http"
	"wifi_kost_be/helper"
	"wifi_kost_be/modules/package/service"

	"github.com/gofiber/fiber/v2"
)

type PackageHandler struct {
	service service.PackageService
}

func NewPackageHandler(service service.PackageService) *PackageHandler {
	return &PackageHandler{service: service}
}

func (h *PackageHandler) GetPackage(c *fiber.Ctx) error {
	type Request struct {
		IsManagedService bool `json:"is_managed_service"`
	}

	// Parse the JSON request
	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid JSON request", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Get the is_managed_service from the request
	is_managed_service := req.IsManagedService
	guestHouses, err := h.service.GetPackage(c.Context(), is_managed_service)
	if err != nil {
		// Handle the error
		response := helper.APIResponse("Failed to retrieve packages", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if guestHouses == nil {
		// Handle case where no packages are found
		response := helper.APIResponse("No packages found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Return the packages
	response := helper.APIResponse("Packages retrieved successfully", http.StatusOK, "OK", guestHouses)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *PackageHandler) GetPackageDetail(c *fiber.Ctx) error {
	type Request struct {
		PackageID int `json:"package_id"`
	}

	// Parse the JSON request
	var req Request
	if err := c.BodyParser(&req); err != nil {
		response := helper.APIResponse("Invalid JSON request", http.StatusBadRequest, "Error", nil)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Get the package_id from the request
	package_id := req.PackageID
	guestHouses, err := h.service.GetPackageDetail(c.Context(), package_id)

	if err != nil {
		// Handle the error
		response := helper.APIResponse("Failed to retrieve packages detail", http.StatusInternalServerError, "Error", nil)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	if guestHouses == nil {
		// Handle case where no packages are found
		response := helper.APIResponse("No package found", http.StatusNotFound, "Error", nil)
		return c.Status(http.StatusNotFound).JSON(response)
	}

	// Return the packages
	response := helper.APIResponse("Package Detail retrieved successfully", http.StatusOK, "OK", guestHouses)
	return c.Status(http.StatusOK).JSON(response)
}
