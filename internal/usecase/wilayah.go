package usecase

import (
	"context"

	"github.com/rahmadfani/wilayah/internal/domain/wilayah"
	"github.com/rahmadfani/wilayah/internal/model"
)

type WilayahUseCase struct {
	wilayahRepo wilayah.Repository
}

func NewWilayahUseCase(wilayahRepo wilayah.Repository) *WilayahUseCase {
	return &WilayahUseCase{
		wilayahRepo: wilayahRepo,
	}
}

func (u *WilayahUseCase) GetProvince(ctx context.Context, query string) ([]model.Wilayah, error) {
	return u.wilayahRepo.GetProvince(ctx, query)
}

func (u *WilayahUseCase) GetRegencyCity(ctx context.Context, code string) ([]model.Wilayah, error) {
	return u.wilayahRepo.GetRegencyCity(ctx, code)
}

func (u *WilayahUseCase) GetDistrict(ctx context.Context, code string) ([]model.Wilayah, error) {
	return u.wilayahRepo.GetDistrict(ctx, code)
}
func (u *WilayahUseCase) GetSubDistrict(ctx context.Context, code string) ([]model.Wilayah, error) {
	return u.wilayahRepo.GetSubDistrict(ctx, code)
}
func (u *WilayahUseCase) GetPostalcode(ctx context.Context, code string) ([]model.WilayahKodepos, error) {
	return u.wilayahRepo.GetPostalcode(ctx, code)
}

func (u *WilayahUseCase) GetWilayah(ctx context.Context) ([]model.Wilayah, error) {
	return u.wilayahRepo.GetWilayah(ctx)
}
