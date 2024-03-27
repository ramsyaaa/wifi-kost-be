package models

import "time"

type UserSubscription struct {
	ID        int64     `json:"id"`
	UserID    int       `json:"user_id"`
	PackageID int       `json:"package_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserSubscription) TableName() string {
	return "user_subscriptions"
}
