package service

import (
	"context"

	"wifi_kost_be/models"
	"wifi_kost_be/modules/guesthouse/repository"
)

type service struct {
	repo repository.GuestHouseRepository
}

func NewGuestHouseService(repo repository.GuestHouseRepository) GuestHouseService {
	return &service{repo: repo}
}

func (s *service) GetGuestHouse(ctx context.Context) ([]*models.GuestHouse, error) {
	return s.repo.GetGuestHouse(ctx)
}

func (s *service) GetGuestHouseDetail(ctx context.Context, guest_house_id int) (*models.GuestHouse, error) {
	return s.repo.GetGuestHouseDetail(ctx, guest_house_id)
}
