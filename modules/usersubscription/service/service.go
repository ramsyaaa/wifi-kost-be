package service

import (
	"context"

	"wifi_kost_be/models"
)

type UserSubscriptionService interface {
	GetUserSubscription(ctx context.Context, user_id int, guest_house_id int) (*models.UserSubscription, error)
}
