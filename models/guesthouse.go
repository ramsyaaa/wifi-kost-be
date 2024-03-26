package models

import "time"

type GuestHouse struct {
	ID               int64     `json:"id"`
	GuestHouseName   string    `json:"guest_house_name"`
	OwnerUserID      int       `json:"owner_user_id"`
	IsManagedService bool      `json:"is_managed_service"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (GuestHouse) TableName() string {
	return "guest_houses"
}
