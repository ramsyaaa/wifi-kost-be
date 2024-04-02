package repository

import (
	"context"
	"errors"
	"time"

	"wifi_kost_be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &repository{db: db}
}

func (r *repository) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *repository) FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("msisdn = ?", msisdn).First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func (r *repository) FindById(ctx context.Context, user_id int) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("id = ?", user_id).First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	return err
}

func (r *repository) GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error) {
	var packages *models.Package
	err := r.db.WithContext(ctx).Where("id = ?", package_id).First(&packages).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return packages, err
}

func (r *repository) GetGuestHouseDetail(ctx context.Context, guest_house_id int) (*models.GuestHouse, error) {
	var guestHouse *models.GuestHouse
	err := r.db.WithContext(ctx).Where("id = ?", guest_house_id).First(&guestHouse).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return guestHouse, err
}

func (r *repository) CreateSubscription(ctx context.Context, subscription *models.UserSubscription) error {
	subscription.StartDate = time.Now().UTC()
	subscription.EndDate = subscription.StartDate.AddDate(0, 1, 0) // Adding 30 days
	subscription.IsActive = true

	err := r.db.WithContext(ctx).Create(subscription).Error
	if err != nil {
		return err
	}

	return nil
}
