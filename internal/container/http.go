package container

import (
	"github.com/rahmadfani/wilayah/internal/handler/http"
	"github.com/rahmadfani/wilayah/internal/repository/mysql"
	"github.com/rahmadfani/wilayah/internal/usecase"
	"github.com/uptrace/bun"
)

type HttpContainer struct {
	WilayahHandler *http.WilayahHandler
}

func NewHttpContainer(db *bun.DB) *HttpContainer {

	container := &HttpContainer{}

	wilayahRepo := mysql.NewWilayahRepository(db)
	wilayahUC := usecase.NewWilayahUseCase(wilayahRepo)
	container.WilayahHandler = http.NewWilayahHandler(wilayahUC)

	return container

}
