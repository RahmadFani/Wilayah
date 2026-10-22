package model

import (
	"github.com/uptrace/bun"
)

type Wilayah struct {
	bun.BaseModel `bun:"table:wilayah"`

	Kode string `bun:",pk"`
	Nama string `bun:",notnull"`
}
