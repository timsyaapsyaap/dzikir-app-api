package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type GeocodeController interface {
	ReverseGeocode(context *gin.Context)
}

type geocodeController struct {
	geocodeService service.GeocodeService
}

func NewGeocodeController(geocodeService service.GeocodeService) GeocodeController {
	return &geocodeController{
		geocodeService: geocodeService,
	}
}

func (controller *geocodeController) ReverseGeocode(context *gin.Context) {
	lat := context.Param("lat")
	lng := context.Param("lng")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Latitude is required", nil)
		context.JSON(400, res)
		return
	}

	longitude, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Longitude is required", nil)
		context.JSON(400, res)
		return
	}

	data, err := controller.geocodeService.ReverseGeocode(latitude, longitude)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}
