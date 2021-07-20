package service

import (
	"log"

	"github.com/afiffahreza/dorayaki-backend/dto"
	"github.com/afiffahreza/dorayaki-backend/entity"
	"github.com/afiffahreza/dorayaki-backend/repository"
	"github.com/mashingan/smapping"
)

type StockService interface {
	Insert(s dto.StockCreateDTO) entity.Stock
	Update(s dto.StockUpdateDTO) entity.Stock
	Delete(s entity.Stock)
	All() []entity.Stock
}

type stockService struct {
	stockRepository repository.StockRepository
}

func NewStockService(stockRepo repository.StockRepository) StockService {
	return &stockService{
		stockRepository: stockRepo,
	}
}

func (service *stockService) Insert(s dto.StockCreateDTO) entity.Stock {
	stock := entity.Stock{}
	err := smapping.FillStruct(&stock, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.stockRepository.InsertStock(stock)
	return res
}

func (service *stockService) Update(s dto.StockUpdateDTO) entity.Stock {
	stock := entity.Stock{}
	err := smapping.FillStruct(&stock, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.stockRepository.UpdateStock(stock)
	return res
}

func (service *stockService) Delete(s entity.Stock) {
	service.stockRepository.DeleteStock(s)
}

func (service *stockService) All() []entity.Stock {
	return service.stockRepository.AllStock()
}
