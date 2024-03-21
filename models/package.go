package models

import "time"

type Package struct {
	ID             int64     `json:"id"`
	PackageName    string    `json:"package_name"`
	PackageSpeed   int       `json:"package_speed"`
	MaximumDevices int       `json:"maximum_devices"`
	ExpiryDay      int       `json:"expiry_day"`
	PackagePrice   int       `json:"package_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Package) TableName() string {
	return "packages"
}
