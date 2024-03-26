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

func NewGuestHouseRepository(db *gorm.DB) GuestHouseRepository {
	return &repository{db: db}
}

func (r *repository) GetGuestHouse(ctx context.Context) ([]*models.GuestHouse, error) {
	var guestHouses []*models.GuestHouse
	err := r.db.WithContext(ctx).Find(&guestHouses).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return guestHouses, err
}

func (r *repository) GetGuestHouseDetail(ctx context.Context, guest_house_id int) (*models.GuestHouse, error) {
	var guestHouse *models.GuestHouse
	err := r.db.WithContext(ctx).Where("id = ?", guest_house_id).First(&guestHouse).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return guestHouse, err
}
