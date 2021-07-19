package dto

type DorayakiCreateDTO struct {
	Rasa      string `json:"rasa" form:"rasa" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" form:"gambar" binding:"required"`
}

type DorayakiUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Rasa      string `json:"rasa" form:"rasa" binding:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" form:"gambar" binding:"required"`
}
