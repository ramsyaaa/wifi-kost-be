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

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) FindByMsisdn(ctx context.Context, msisdn string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("msisdn = ?", msisdn).First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}
