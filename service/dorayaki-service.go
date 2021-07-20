package service

import (
	"log"

	"github.com/afiffahreza/dorayaki-backend/dto"
	"github.com/afiffahreza/dorayaki-backend/entity"
	"github.com/afiffahreza/dorayaki-backend/repository"
	"github.com/mashingan/smapping"
)

type DorayakiService interface {
	Insert(d dto.DorayakiCreateDTO) entity.Dorayaki
	Update(d dto.DorayakiUpdateDTO) entity.Dorayaki
	Delete(d entity.Dorayaki)
	All() []entity.Dorayaki
	FindByID(id uint64) entity.Dorayaki
}

type dorayakiService struct {
	dorayakiRepository repository.DorayakiRepository
}

func NewDorayakiService(dorayakiRepo repository.DorayakiRepository) DorayakiService {
	return &dorayakiService{
		dorayakiRepository: dorayakiRepo,
	}
}

func (service *dorayakiService) Insert(d dto.DorayakiCreateDTO) entity.Dorayaki {
	dorayaki := entity.Dorayaki{}
	err := smapping.FillStruct(&dorayaki, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.dorayakiRepository.InsertDorayaki(dorayaki)
	return res
}

func (service *dorayakiService) Update(d dto.DorayakiUpdateDTO) entity.Dorayaki {
	dorayaki := entity.Dorayaki{}
	err := smapping.FillStruct(&dorayaki, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.dorayakiRepository.UpdateDorayaki(dorayaki)
	return res
}

func (service *dorayakiService) Delete(d entity.Dorayaki) {
	service.dorayakiRepository.DeleteDorayaki(d)
}

func (service *dorayakiService) All() []entity.Dorayaki {
	return service.dorayakiRepository.AllDorayaki()
}

func (service *dorayakiService) FindByID(id uint64) entity.Dorayaki {
	return service.dorayakiRepository.FindByID(id)
}
