package service

import (
	"context"

	"wifi_kost_be/models"
)

type GuestHouseService interface {
	GetGuestHouse(ctx context.Context) ([]*models.GuestHouse, error)
}
