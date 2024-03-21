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

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &repository{db: db}
}

func (r *repository) GetPackage(ctx context.Context) ([]*models.Package, error) {
	var packages []*models.Package
	err := r.db.WithContext(ctx).Find(&packages).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return packages, err
}

func (r *repository) GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error) {
	var packages *models.Package
	err := r.db.WithContext(ctx).Where("id = ?", package_id).First(&packages).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return packages, err
}
