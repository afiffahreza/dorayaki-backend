package controller

import (
	"net/http"
	"strconv"

	"github.com/afiffahreza/dorayaki-backend/dto"
	"github.com/afiffahreza/dorayaki-backend/entity"
	"github.com/afiffahreza/dorayaki-backend/helper"
	"github.com/afiffahreza/dorayaki-backend/service"
	"github.com/gin-gonic/gin"
)

type StockController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindStockByStoreID(ctx *gin.Context)
}

type stockController struct {
	stockService service.StockService
}

func NewStockController(stockServ service.StockService) StockController {
	return &stockController{
		stockService: stockServ,
	}
}

func (c *stockController) Insert(ctx *gin.Context) {
	var stockCreateDTO dto.StockCreateDTO
	err := ctx.ShouldBind(&stockCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.stockService.Insert(stockCreateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *stockController) Update(ctx *gin.Context) {
	var stockUpdateDTO dto.StockUpdateDTO
	err := ctx.ShouldBind(&stockUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.stockService.Update(stockUpdateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *stockController) Delete(ctx *gin.Context) {
	var stock entity.Stock
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		stock.ID = id
		c.stockService.Delete(stock)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *stockController) All(ctx *gin.Context) {
	var stocks []entity.Stock = c.stockService.All()
	res := helper.BuildResponse(true, "OK", stocks)
	ctx.JSON(http.StatusOK, res)
}

func (c *stockController) FindStockByStoreID(ctx *gin.Context) {
	var stocks []entity.Stock
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		stocks = c.stockService.FindStockByStoreID(id)
		res := helper.BuildResponse(true, "OK", stocks)
		ctx.JSON(http.StatusOK, res)
	}
}
