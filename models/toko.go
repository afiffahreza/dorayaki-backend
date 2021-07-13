package models

type toko struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nama      string `json:"nama"`
	Jalan     string `json:"alamat"`
	Kecamatan string `json:"kecamatan"`
	Provinsi  string `json:"provinsi"`
}
