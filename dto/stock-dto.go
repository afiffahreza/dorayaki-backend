package dto

type StockCreateDTO struct {
	TokoID     uint64 `json:"tokoid" form:"tokoid" binding:"required"`
	DorayakiID uint64 `json:"dorayakiid" form:"dorayakiid" binding:"required"`
	Quantity   uint64 `json:"quantity" form:"quantity" binding:"required"`
}

type StockUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	TokoID     uint64 `json:"tokoid" form:"tokoid" binding:"required"`
	DorayakiID uint64 `json:"dorayakiid" form:"dorayakiid" binding:"required"`
	Quantity   uint64 `json:"quantity" form:"quantity" binding:"required"`
}
