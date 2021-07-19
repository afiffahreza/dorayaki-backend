package entity

type Dorayaki struct {
	ID        uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Rasa      string `json:"rasa"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
}
