package models

import "time"

type Transaction struct {
	ID                int64     `json:"id"`
	UserID            int       `json:"user_id"`
	PackageID         int       `json:"package_id"`
	GuestHouseID      int       `json:"guest_house_id"`
	Amount            int       `json:"amount"`
	TotalAmount       int       `json:"total_amount"`
	IsUsingCoupon     bool      `json:"is_using_coupon"`
	PaymentMethod     string    `json:"payment_method"`
	TransactionStatus string    `json:"transaction_status"`
	TransactionDate   time.Time `json:"transaction_date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
