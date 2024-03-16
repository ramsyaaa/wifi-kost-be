package service

import (
	"context"
	"errors"

	"wifi_kost_be/models"
)

type UserService interface {
	FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error)
}

type TokenService interface {
	GenerateToken(user *models.User) (string, error)
	GenerateGuestToken(msisdn string) (string, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

type LoginData struct {
	Msisdn   string
	Password string
}
