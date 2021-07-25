package connection

import (
	"log"

	"github.com/afiffahreza/dorayaki-backend/entity"
	"gorm.io/gorm"
)

var tokos = []entity.Toko{
	{
		ID:        1,
		Nama:      "Dorayaki Bros",
		Jalan:     "Jl. Raya Ragunan No. 13",
		Kecamatan: "Pasar Minggu",
		Provinsi:  "DKI Jakarta",
	},
	{
		ID:        2,
		Nama:      "Mc Dorayaki's",
		Jalan:     "Jl. TB Simatupang No. 3",
		Kecamatan: "Jagakarsa",
		Provinsi:  "DKI Jakarta",
	},
	{
		ID:        3,
		Nama:      "Dorayaki Geprek Crisbar",
		Jalan:     "Jl. Cisitu Indah No. 2",
		Kecamatan: "Coblong",
		Provinsi:  "Jawa Barat",
	},
}

var dorayakis = []entity.Dorayaki{
	{
		ID:        1,
		Rasa:      "Cokelat",
		Deskripsi: "Cokelat dari Swiss",
		Gambar:    "https://res.cloudinary.com/cloudinary-afif/image/upload/v1627226682/d1_tnckhu.jpg",
	},
}

var stocks = []entity.Stock{
	{
		ID:         1,
		TokoID:     1,
		DorayakiID: 1,
		Quantity:   30,
	},
	{
		ID:         2,
		TokoID:     3,
		DorayakiID: 1,
		Quantity:   50,
	},
}

func LoadSeed(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&entity.Toko{}, &entity.Dorayaki{}, &entity.Stock{})
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	db.AutoMigrate(&entity.Toko{}, &entity.Dorayaki{}, &entity.Stock{})

	for i := range tokos {
		err = db.Debug().Model(&entity.Toko{}).Create(&tokos[i]).Error
		if err != nil {
			log.Fatalf("cannot seed tokos table: %v", err)
		}
	}

	for i := range dorayakis {
		err = db.Debug().Model(&entity.Dorayaki{}).Create(&dorayakis[i]).Error
		if err != nil {
			log.Fatalf("cannot seed dorayakis table: %v", err)
		}
	}

	for i := range stocks {
		err = db.Debug().Model(&entity.Stock{}).Create(&stocks[i]).Error
		if err != nil {
			log.Fatalf("cannot seed stocks table: %v", err)
		}
	}
}
