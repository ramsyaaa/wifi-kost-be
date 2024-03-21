package service

import (
	"context"

	"wifi_kost_be/models"
	"wifi_kost_be/modules/package/repository"
)

type service struct {
	repo repository.PackageRepository
}

func NewPackageService(repo repository.PackageRepository) PackageService {
	return &service{repo: repo}
}

func (s *service) GetPackage(ctx context.Context) ([]*models.Package, error) {
	return s.repo.GetPackage(ctx)
}

func (s *service) GetPackageDetail(ctx context.Context, package_id int) (*models.Package, error) {
	return s.repo.GetPackageDetail(ctx, package_id)
}
