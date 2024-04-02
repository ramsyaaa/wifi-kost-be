package service

import (
	"context"

	"wifi_kost_be/models"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error)
	FindById(ctx context.Context, user_id int) (*models.User, error)
	CreateUser(ctx context.Context, msisdn string) (*models.User, error)
	GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error)
	GetGuestHouseDetail(ctx context.Context, guest_house_id int) (*models.GuestHouse, error)
	CreateSubscription(ctx context.Context, subscription *models.UserSubscription) error
}
