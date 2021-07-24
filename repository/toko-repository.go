package repository

import (
	"github.com/afiffahreza/dorayaki-backend/entity"
	"gorm.io/gorm"
)

type TokoRepository interface {
	InsertToko(t entity.Toko) entity.Toko
	UpdateToko(t entity.Toko) entity.Toko
	DeleteToko(t entity.Toko)
	AllToko() []entity.Toko
	FindTokoQuery(query string) []entity.Toko
	FindByID(id uint64) entity.Toko
}

type tokoConnection struct {
	connection *gorm.DB
}

func NewTokoRepository(dbConn *gorm.DB) TokoRepository {
	return &tokoConnection{
		connection: dbConn,
	}
}

func (db *tokoConnection) InsertToko(t entity.Toko) entity.Toko {
	db.connection.Save(&t)
	return t
}

func (db *tokoConnection) UpdateToko(t entity.Toko) entity.Toko {
	db.connection.Save(&t)
	return t
}

func (db *tokoConnection) DeleteToko(t entity.Toko) {
	db.connection.Delete(&t)
}

func (db *tokoConnection) AllToko() []entity.Toko {
	var tokoAll []entity.Toko
	db.connection.Find(&tokoAll)
	return tokoAll
}

func (db *tokoConnection) FindByID(id uint64) entity.Toko {
	var toko entity.Toko
	db.connection.Find(&toko, id)
	return toko
}

func (db *tokoConnection) FindTokoQuery(query string) []entity.Toko {
	var tokoQuery []entity.Toko
	finalQuery := "%" + query + "%"
	db.connection.Where("nama LIKE ? OR kecamatan LIKE ? OR provinsi LIKE ?", finalQuery, finalQuery, finalQuery).Find(&tokoQuery)
	return tokoQuery
}
