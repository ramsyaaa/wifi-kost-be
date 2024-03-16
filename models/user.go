package models

import "time"

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Msisdn      string    `json:"msisdn"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
