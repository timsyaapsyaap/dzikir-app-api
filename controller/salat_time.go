package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type SalatTimeController interface {
	FindCity(context *gin.Context)
	CityDetails(context *gin.Context)
	AllCities(context *gin.Context)
	Schedule(context *gin.Context)
}

type salatTimeController struct {
	salatTimeService service.SalatTimeService
}

func NewSalatTimeController(salatTimeService service.SalatTimeService) SalatTimeController {
	return &salatTimeController{
		salatTimeService: salatTimeService,
	}
}

func (controller *salatTimeController) FindCity(context *gin.Context) {
	data, err := controller.salatTimeService.FindCity(context.Param("city"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) CityDetails(context *gin.Context) {
	data, err := controller.salatTimeService.CityDetails(context.Param("id"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) AllCities(context *gin.Context) {
	data, err := controller.salatTimeService.AllCities()
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) Schedule(context *gin.Context) {
	cityId, err := strconv.Atoi(context.Param("cityId"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	year, err := strconv.Atoi(context.Param("year"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	month, err := strconv.Atoi(context.Param("month"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	date, err := strconv.Atoi(context.Param("date"))
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	data, err := controller.salatTimeService.Schedule(cityId, year, month, date)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}
