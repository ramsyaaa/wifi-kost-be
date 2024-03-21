package repository

import (
	"context"

	"wifi_kost_be/models"
)

type PackageRepository interface {
	GetPackage(ctx context.Context) ([]*models.Package, error)
	GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error)
}
