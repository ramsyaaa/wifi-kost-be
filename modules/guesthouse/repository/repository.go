package repository

import (
	"context"

	"wifi_kost_be/models"
)

type GuestHouseRepository interface {
	GetGuestHouse(ctx context.Context) ([]*models.GuestHouse, error)
}
