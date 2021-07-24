package dto

type TokoCreateDTO struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Jalan     string `json:"jalan" form:"nama" binding:"required"`
	Kecamatan string `json:"kecamatan" form:"kecamatan" binding:"required"`
	Provinsi  string `json:"provinsi" form:"provinsi" binding:"required"`
}

type TokoUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Jalan     string `json:"jalan" form:"nama" binding:"required"`
	Kecamatan string `json:"kecamatan" form:"kecamatan" binding:"required"`
	Provinsi  string `json:"provinsi" form:"provinsi" binding:"required"`
}
