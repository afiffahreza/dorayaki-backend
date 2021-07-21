package repository

import (
	"github.com/afiffahreza/dorayaki-backend/entity"
	"gorm.io/gorm"
)

type StockRepository interface {
	InsertStock(s entity.Stock) entity.Stock
	UpdateStock(s entity.Stock) entity.Stock
	DeleteStock(s entity.Stock)
	AllStock() []entity.Stock
	FindStockByStoreID(id uint64) []entity.Stock
}

type stockConnection struct {
	connection *gorm.DB
}

func NewStockRepository(dbConn *gorm.DB) StockRepository {
	return &stockConnection{
		connection: dbConn,
	}
}

func (db *stockConnection) InsertStock(s entity.Stock) entity.Stock {
	db.connection.Save(&s)
	return s
}

func (db *stockConnection) UpdateStock(s entity.Stock) entity.Stock {
	db.connection.Save(&s)
	return s
}

func (db *stockConnection) DeleteStock(s entity.Stock) {
	db.connection.Delete(&s)
}

func (db *stockConnection) AllStock() []entity.Stock {
	var stocks []entity.Stock
	db.connection.Find(&stocks)
	return stocks
}

func (db *stockConnection) FindStockByStoreID(id uint64) []entity.Stock {
	var stocks []entity.Stock
	db.connection.Where("toko_id = ?", id).Find(&stocks)
	return stocks
}
