package entity

type Toko struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Nama      string `json:"nama"`
	Jalan     string `json:"jalan"`
	Kecamatan string `json:"kecamatan"`
	Provinsi  string `json:"provinsi"`
}
