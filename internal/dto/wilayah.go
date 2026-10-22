package dto

type WilayahResponse struct {
	Code       string `json:"code"`
	Nama       string `json:"nama"`
	ParentCode string `json:"parent_code"`
	Tingkat    string `json:"tingkat"`
}
