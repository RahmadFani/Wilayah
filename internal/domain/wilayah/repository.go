package wilayah

import (
	"context"

	"github.com/rahmadfani/wilayah/internal/model"
)

type Repository interface {
	GetWilayah(ctx context.Context) ([]model.Wilayah, error)
	GetProvince(ctx context.Context) ([]model.Wilayah, error)
	GetRegencyCity(ctx context.Context, code string) ([]model.Wilayah, error)
	GetDistrict(ctx context.Context, code string) ([]model.Wilayah, error)
	GetSubDistrict(ctx context.Context, code string) ([]model.Wilayah, error)
	GetPostalcode(ctx context.Context, code string) ([]model.WilayahKodepos, error)
}
