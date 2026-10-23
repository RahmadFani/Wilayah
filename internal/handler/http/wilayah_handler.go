package http

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/rahmadfani/wilayah/internal/dto"
	"github.com/rahmadfani/wilayah/internal/usecase"
)

type WilayahHandler struct {
	wilayahUC *usecase.WilayahUseCase
}

func NewWilayahHandler(wilayahUC *usecase.WilayahUseCase) *WilayahHandler {
	return &WilayahHandler{
		wilayahUC: wilayahUC,
	}
}

func (h *WilayahHandler) PublicRouteWilayah(route fiber.Router) {
	route.Get("/province", h.Porvince)
	route.Get("/city", h.City)
	route.Get("/district", h.District)
	route.Get("/subdistrict", h.SubDistrict)
	route.Get("/postalcode", h.Postalcode)
}

func (h *WilayahHandler) Porvince(c fiber.Ctx) error {

	wilayah, err := h.wilayahUC.GetProvince(c.RequestCtx())
	if err != nil {
		return HttpError(c, http.StatusNotFound, err.Error())
	}

	res := []dto.WilayahResponse{}
	for _, w := range wilayah {
		res = append(res, dto.WilayahResponse{
			Code:       w.Kode,
			Nama:       w.Nama,
			ParentCode: "",
			Tingkat:    "Provinsi",
		})
	}

	return HttpSuccess(c, http.StatusAccepted, res)
}

func (h *WilayahHandler) City(c fiber.Ctx) error {

	code := c.Query("code")

	wilayah, err := h.wilayahUC.GetRegencyCity(c.RequestCtx(), code)
	if err != nil {
		return HttpError(c, http.StatusNotFound, err.Error())
	}

	res := []dto.WilayahResponse{}
	for _, w := range wilayah {
		res = append(res, dto.WilayahResponse{
			Code:       w.Kode,
			Nama:       w.Nama,
			ParentCode: w.Kode[:2],
			Tingkat:    "Kabupaten/Kota",
		})
	}

	return HttpSuccess(c, http.StatusAccepted, res)
}

func (h *WilayahHandler) District(c fiber.Ctx) error {

	code := c.Query("code")

	wilayah, err := h.wilayahUC.GetDistrict(c.RequestCtx(), code)
	if err != nil {
		return HttpError(c, http.StatusNotFound, err.Error())
	}

	res := []dto.WilayahResponse{}
	for _, w := range wilayah {
		res = append(res, dto.WilayahResponse{
			Code:       w.Kode,
			Nama:       w.Nama,
			ParentCode: w.Kode[:5],
			Tingkat:    "Kecamatan",
		})
	}

	return HttpSuccess(c, http.StatusAccepted, res)
}

func (h *WilayahHandler) SubDistrict(c fiber.Ctx) error {

	code := c.Query("code")

	wilayah, err := h.wilayahUC.GetSubDistrict(c.RequestCtx(), code)
	if err != nil {
		return HttpError(c, http.StatusNotFound, err.Error())
	}

	res := []dto.WilayahResponse{}
	for _, w := range wilayah {
		res = append(res, dto.WilayahResponse{
			Code:       w.Kode,
			Nama:       w.Nama,
			ParentCode: w.Kode[:8],
			Tingkat:    "Desa",
		})
	}

	return HttpSuccess(c, http.StatusAccepted, res)
}

func (h *WilayahHandler) Postalcode(c fiber.Ctx) error {

	code := c.Query("code")

	kodepos, err := h.wilayahUC.GetPostalcode(c.RequestCtx(), code)
	if err != nil {
		return HttpError(c, http.StatusNotFound, err.Error())
	}

	return HttpSuccess(c, http.StatusAccepted, kodepos)
}
