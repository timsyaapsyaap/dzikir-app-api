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
	city := context.Param("city")
	if city == "" {
		res := helper.BuildErrorResponse("Bad Request", "City is required", nil)
		context.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.FindCity(city)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) CityDetails(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		res := helper.BuildErrorResponse("Bad Request", "City ID is required", nil)
		context.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.CityDetails(id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) AllCities(context *gin.Context) {
	data, err := controller.salatTimeService.AllCities()
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *salatTimeController) Schedule(context *gin.Context) {
	cityId, err := strconv.Atoi(context.Param("cityId"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "City ID is required or not valid", nil)
		context.JSON(400, res)
		return
	}

	year, err := strconv.Atoi(context.Param("year"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Year is required or not valid", nil)
		context.JSON(400, res)
		return
	}

	month, err := strconv.Atoi(context.Param("month"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Month is required or not valid", nil)
		context.JSON(400, res)
		return
	}

	date, err := strconv.Atoi(context.Param("date"))
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Date is required or not valid", nil)
		context.JSON(400, res)
		return
	}

	data, err := controller.salatTimeService.Schedule(cityId, year, month, date)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}
