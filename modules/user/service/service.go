package service

import (
	"context"
	"errors"

	"wifi_kost_be/models"
)

type UserService interface {
	FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error)
	CreateUser(ctx context.Context, msisdn string) (*models.User, error)
}

type TokenService interface {
	GenerateToken(user *models.User, guest_house_id int) (string, error)
	GenerateGuestToken(msisdn string, guest_house_id int) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

type LoginData struct {
	Msisdn   string
	Password string
}
