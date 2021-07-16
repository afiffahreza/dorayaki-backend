package service

import (
	"log"

	"github.com/afiffahreza/dorayaki-backend/dto"
	"github.com/afiffahreza/dorayaki-backend/entity"
	"github.com/afiffahreza/dorayaki-backend/repository"
	"github.com/mashingan/smapping"
)

type TokoService interface {
	Insert(t dto.TokoCreateDTO) entity.Toko
	Update(t dto.TokoUpdateDTO) entity.Toko
	Delete(t entity.Toko)
	All() []entity.Toko
	FindTokoKecamatan(kecamatan string) []entity.Toko
	FindTokoProvinsi(provinsi string) []entity.Toko
}

type tokoService struct {
	tokoRepository repository.TokoRepository
}

func NewTokoService(tokoRepo repository.TokoRepository) TokoService {
	return &tokoService{
		tokoRepository: tokoRepo,
	}
}

func (service *tokoService) Insert(t dto.TokoCreateDTO) entity.Toko {
	toko := entity.Toko{}
	err := smapping.FillStruct(&toko, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.tokoRepository.InsertToko(toko)
	return res
}

func (service *tokoService) Update(t dto.TokoUpdateDTO) entity.Toko {
	toko := entity.Toko{}
	err := smapping.FillStruct(&toko, smapping.MapFields(&t))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.tokoRepository.UpdateToko(toko)
	return res
}

func (service *tokoService) Delete(t entity.Toko) {
	service.tokoRepository.DeleteToko(t)
}

func (service *tokoService) All() []entity.Toko {
	return service.tokoRepository.AllToko()
}

func (service *tokoService) FindTokoKecamatan(kecamatan string) []entity.Toko {
	return service.tokoRepository.FindTokoKecamatan(kecamatan)
}

func (service *tokoService) FindTokoProvinsi(provinsi string) []entity.Toko {
	return service.tokoRepository.FindTokoProvinsi(provinsi)
}
