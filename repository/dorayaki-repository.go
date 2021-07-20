package repository

import (
	"github.com/afiffahreza/dorayaki-backend/entity"
	"gorm.io/gorm"
)

type DorayakiRepository interface {
	InsertDorayaki(d entity.Dorayaki) entity.Dorayaki
	UpdateDorayaki(d entity.Dorayaki) entity.Dorayaki
	DeleteDorayaki(d entity.Dorayaki)
	AllDorayaki() []entity.Dorayaki
	FindByID(id uint64) entity.Dorayaki
}

type dorayakiConnection struct {
	connection *gorm.DB
}

func NewDorayakiRepository(dbConn *gorm.DB) DorayakiRepository {
	return &dorayakiConnection{
		connection: dbConn,
	}
}

func (db *dorayakiConnection) InsertDorayaki(d entity.Dorayaki) entity.Dorayaki {
	db.connection.Save(&d)
	return d
}

func (db *dorayakiConnection) UpdateDorayaki(d entity.Dorayaki) entity.Dorayaki {
	db.connection.Save(&d)
	return d
}

func (db *dorayakiConnection) DeleteDorayaki(d entity.Dorayaki) {
	db.connection.Delete(&d)
}

func (db *dorayakiConnection) AllDorayaki() []entity.Dorayaki {
	var dorayakiAll []entity.Dorayaki
	db.connection.Find(&dorayakiAll)
	return dorayakiAll
}

func (db *dorayakiConnection) FindByID(id uint64) entity.Dorayaki {
	var dorayaki entity.Dorayaki
	db.connection.Find(&dorayaki, id)
	return dorayaki
}
