package model

import "github.com/uptrace/bun"

type WilayahKodepos struct {
	bun.BaseModel `bun:"table:wilayah_kodepos"`

	Kode    string `bun:",pk" json:"kode"`
	Kodepos string `bun:",notnull" json:"kodepos"`
}
