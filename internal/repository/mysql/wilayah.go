package mysql

import (
	"context"
	"log/slog"

	"github.com/rahmadfani/wilayah/internal/domain/wilayah"
	"github.com/rahmadfani/wilayah/internal/model"
	"github.com/uptrace/bun"
)

type wilayahRepository struct {
	db *bun.DB
}

// GetWilayah implements wilayah.Repository.
func (w *wilayahRepository) GetWilayah(ctx context.Context) ([]model.Wilayah, error) {
	var wilayahs []model.Wilayah

	if err := w.db.NewSelect().
		Model(&wilayahs).
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return wilayahs, nil
}

// GetPostalcode implements wilayah.Repository.
func (w *wilayahRepository) GetPostalcode(ctx context.Context, code string) ([]model.WilayahKodepos, error) {
	var kodepos []model.WilayahKodepos
	if err := w.db.NewSelect().
		Model(&kodepos).
		Where("kode LIKE ?", code+"%").
		Order("kodepos ASC").
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return kodepos, nil
}

// GetSubDistrict implements wilayah.Repository.
func (w *wilayahRepository) GetSubDistrict(ctx context.Context, code string) ([]model.Wilayah, error) {
	var wilayahs []model.Wilayah
	if err := w.db.NewSelect().
		Model(&wilayahs).
		Where("kode LIKE ?", code+".%").
		Where("LENGTH(kode) = 13").
		Order("nama ASC").
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return wilayahs, nil
}

// GetDistrict implements wilayah.Repository.
func (w *wilayahRepository) GetDistrict(ctx context.Context, code string) ([]model.Wilayah, error) {
	var wilayahs []model.Wilayah
	if err := w.db.NewSelect().
		Model(&wilayahs).
		Where("kode LIKE ?", code+".%").
		Where("LENGTH(kode) = 8").
		Order("nama ASC").
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return wilayahs, nil
}

// GetRegencyCity implements wilayah.Repository.
func (w *wilayahRepository) GetRegencyCity(ctx context.Context, code string) ([]model.Wilayah, error) {
	var wilayahs []model.Wilayah
	if err := w.db.NewSelect().
		Model(&wilayahs).
		Where("kode LIKE ?", code+".%").
		Where("LENGTH(kode) = 5").
		Order("nama ASC").
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return wilayahs, nil
}

// GetProvince implements wilayah.Repository.
func (w *wilayahRepository) GetProvince(ctx context.Context) ([]model.Wilayah, error) {
	var wilayahs []model.Wilayah
	if err := w.db.NewSelect().
		Model(&wilayahs).
		Where("kode NOT LIKE ?", "%.%.%").
		Where("kode NOT LIKE ?", "%.%").
		Order("nama ASC").
		Scan(ctx); err != nil {
		slog.Error("DB Scan", "error_message", err.Error())
		return nil, err
	}

	return wilayahs, nil
}

func NewWilayahRepository(db *bun.DB) wilayah.Repository {
	return &wilayahRepository{
		db: db,
	}
}
