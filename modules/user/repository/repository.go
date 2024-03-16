package repository

import (
	"context"

	"wifi_kost_be/models"
)

type UserRepository interface {
	FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error)
}
