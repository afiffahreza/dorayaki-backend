package entity

type Stock struct {
	ID         uint64   `json:"id" gorm:"primary_key:auto_increment"`
	TokoID     uint64   `json:"tokoid"`
	DorayakiID uint64   `json:"dorayakiid"`
	Quantity   uint64   `json:"quantity"`
	Toko       Toko     `gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Dorayaki   Dorayaki `gorm:"foreignKey:DorayakiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
