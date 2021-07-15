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
	//FindTokoProvinsi(provinsi string) []entity.Toko
	//FindTokoKecamatan(kecamatan string) []entity.Toko
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
