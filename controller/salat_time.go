package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type SalatTimeController interface {
	FindCity(ctx *gin.Context)
	CityDetails(ctx *gin.Context)
	AllCities(ctx *gin.Context)
	Schedule(ctx *gin.Context)
}

type salatTimeController struct {
	salatTimeService service.SalatTimeService
}

func NewSalatTimeController(salatTimeService service.SalatTimeService) SalatTimeController {
	return &salatTimeController{
		salatTimeService: salatTimeService,
	}
}

func (controller *salatTimeController) FindCity(ctx *gin.Context) {
	city := ctx.Param("city")
	if city == "" {
		res := helper.BuildErrorResponse("Bad Request", "City is required", nil)
		ctx.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.FindCity(ctx, city)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}

func (controller *salatTimeController) CityDetails(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		res := helper.BuildErrorResponse("Bad Request", "City ID is required", nil)
		ctx.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.CityDetails(ctx, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}

func (controller *salatTimeController) AllCities(ctx *gin.Context) {
	data, err := controller.salatTimeService.AllCities(ctx)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}

func (controller *salatTimeController) Schedule(ctx *gin.Context) {
	cityId, err := strconv.Atoi(ctx.Param("cityId"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "City ID is required or not valid", nil)
		ctx.JSON(400, res)
		return
	}

	year, err := strconv.Atoi(ctx.Param("year"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Year is required or not valid", nil)
		ctx.JSON(400, res)
		return
	}

	month, err := strconv.Atoi(ctx.Param("month"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Month is required or not valid", nil)
		ctx.JSON(400, res)
		return
	}

	date, err := strconv.Atoi(ctx.Param("date"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Date is required or not valid", nil)
		ctx.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.Schedule(ctx, cityId, year, month, date)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}
