package service

import (
	"context"
	"errors"
	"time"

	"wifi_kost_be/models"
	"wifi_kost_be/modules/transaction/repository"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &service{repo: repo}
}

func (s *service) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	return s.repo.CreateTransaction(ctx, transaction)
}

func (s *service) FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error) {
	return s.repo.FindByMsisdn(ctx, msisdn)
}
func (s *service) FindById(ctx context.Context, user_id int) (*models.User, error) {
	return s.repo.FindById(ctx, user_id)
}

func (s *service) CreateUser(ctx context.Context, msisdn string) (*models.User, error) {
	// Check if the user already exists
	user, err := s.repo.FindByMsisdn(ctx, msisdn)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("user already exists")
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user with default password
	user = &models.User{
		Msisdn:    msisdn,
		Role:      "user",
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the user to the database
	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error) {
	return s.repo.GetPackageDetail(ctx, package_id)
}

func (s *service) GetGuestHouseDetail(ctx context.Context, guest_house_id int) (*models.GuestHouse, error) {
	return s.repo.GetGuestHouseDetail(ctx, guest_house_id)
}

func (s *service) CreateSubscription(ctx context.Context, subscription *models.UserSubscription) error {
	return s.repo.CreateSubscription(ctx, subscription)
}
