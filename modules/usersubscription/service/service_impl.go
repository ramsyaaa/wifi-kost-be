package service

import (
	"context"

	"wifi_kost_be/models"
	"wifi_kost_be/modules/usersubscription/repository"
)

type service struct {
	repo repository.UserSubscriptionRepository
}

func NewUserSubscriptionService(repo repository.UserSubscriptionRepository) UserSubscriptionService {
	return &service{repo: repo}
}

func (s *service) GetUserSubscription(ctx context.Context, user_id int, guest_house_id int) (*models.UserSubscription, error) {
	return s.repo.GetUserSubscription(ctx, user_id, guest_house_id)
}
