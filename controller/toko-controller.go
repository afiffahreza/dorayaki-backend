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

type TokoController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindTokoQuery(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type tokoController struct {
	tokoService service.TokoService
}

func NewTokoController(tokoServ service.TokoService) TokoController {
	return &tokoController{
		tokoService: tokoServ,
	}
}

func (c *tokoController) Insert(ctx *gin.Context) {
	var tokoCreateDTO dto.TokoCreateDTO
	err := ctx.ShouldBind(&tokoCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.tokoService.Insert(tokoCreateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *tokoController) Update(ctx *gin.Context) {
	var tokoUpdateDTO dto.TokoUpdateDTO
	err := ctx.ShouldBind(&tokoUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.tokoService.Update(tokoUpdateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *tokoController) Delete(ctx *gin.Context) {
	var toko entity.Toko
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		toko.ID = id
		c.tokoService.Delete(toko)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *tokoController) All(ctx *gin.Context) {
	var tokoAll []entity.Toko = c.tokoService.All()
	res := helper.BuildResponse(true, "OK", tokoAll)
	ctx.JSON(http.StatusOK, res)
}

func (c *tokoController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var toko entity.Toko = c.tokoService.FindByID(id)
	if (toko == entity.Toko{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", toko)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *tokoController) FindTokoQuery(ctx *gin.Context) {
	var query string = ctx.DefaultQuery("query", "__EMPTY__")
	var tokoQuery []entity.Toko = c.tokoService.FindTokoQuery(query)
	if len(tokoQuery) == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", tokoQuery)
		ctx.JSON(http.StatusOK, res)
	}
}
