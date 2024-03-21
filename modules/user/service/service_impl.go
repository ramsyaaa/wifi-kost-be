package service

import (
	"context"
	"time"

	"wifi_kost_be/models"
	"wifi_kost_be/modules/user/repository"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &service{repo: repo}
}

type tokenService struct {
	// You can add any dependencies required for token generation here
	secretKey string
}

func NewTokenService(secretKey string) TokenService {
	return &tokenService{secretKey: secretKey}
}

func (s *service) FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error) {
	return s.repo.FindByMsisdn(ctx, msisdn)
}

func (s *tokenService) GenerateToken(user *models.User, guest_house_id int) (string, error) {
	var rid int
	switch user.Role {
	case "owner":
		rid = 1
	case "user":
		rid = 2
	default:
		rid = 0
	}
	// Create a new JWT token with the user's ID as the "sub" claim and set the expiry time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":            user.ID,
		"name":           user.Name,
		"msisdn":         user.Msisdn,
		"guest_house_id": guest_house_id,
		"rid":            rid,
		"exp":            time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *tokenService) GenerateGuestToken(msisdn string, guest_house_id int) (string, error) {
	// Create a new JWT token with the user's ID as the "sub" claim and set the expiry time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":            0,
		"msisdn":         msisdn,
		"guest_house_id": guest_house_id,
		"rid":            0,
		"exp":            time.Now().Add(time.Hour).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) Login(ctx context.Context, msisdn, password string) (*models.User, error) {

	user, err := s.repo.FindByMsisdn(ctx, msisdn)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Password != password {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
