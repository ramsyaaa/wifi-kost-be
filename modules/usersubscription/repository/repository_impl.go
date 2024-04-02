package repository

import (
	"context"
	"errors"

	"wifi_kost_be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserSubscriptionRepository(db *gorm.DB) UserSubscriptionRepository {
	return &repository{db: db}
}

func (r *repository) GetUserSubscription(ctx context.Context, user_id int) (*models.UserSubscription, error) {
	var packages *models.UserSubscription
	err := r.db.WithContext(ctx).Where("user_id = ?", user_id).Where("is_active = ?", true).Last(&packages).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return packages, err
}
