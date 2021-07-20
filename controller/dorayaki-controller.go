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

type DorayakiController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type dorayakiController struct {
	dorayakiService service.DorayakiService
}

func NewDorayakiController(dorayakiServ service.DorayakiService) DorayakiController {
	return &dorayakiController{
		dorayakiService: dorayakiServ,
	}
}

func (c *dorayakiController) Insert(ctx *gin.Context) {
	var dorayakiCreateDTO dto.DorayakiCreateDTO
	err := ctx.ShouldBind(&dorayakiCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.dorayakiService.Insert(dorayakiCreateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *dorayakiController) Update(ctx *gin.Context) {
	var dorayakiUpdateDTO dto.DorayakiUpdateDTO
	err := ctx.ShouldBind(&dorayakiUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		res := c.dorayakiService.Update(dorayakiUpdateDTO)
		response := helper.BuildResponse(true, "OK", res)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *dorayakiController) Delete(ctx *gin.Context) {
	var dorayaki entity.Dorayaki
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		dorayaki.ID = id
		c.dorayakiService.Delete(dorayaki)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *dorayakiController) All(ctx *gin.Context) {
	var dorayakiAll []entity.Dorayaki = c.dorayakiService.All()
	res := helper.BuildResponse(true, "OK", dorayakiAll)
	ctx.JSON(http.StatusOK, res)
}

func (c *dorayakiController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var dorayaki entity.Dorayaki = c.dorayakiService.FindByID(id)
	if (dorayaki == entity.Dorayaki{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", dorayaki)
		ctx.JSON(http.StatusOK, res)
	}
}
