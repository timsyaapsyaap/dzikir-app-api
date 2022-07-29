package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type HijriController interface {
	GregorianToHijri(context *gin.Context)
}

type hijriController struct {
	hijriService service.HijriService
}

func NewHijriController(hijriService service.HijriService) HijriController {
	return &hijriController{
		hijriService: hijriService,
	}
}

func (controller *hijriController) GregorianToHijri(context *gin.Context) {
	date := context.Param("date")
	month := context.Param("month")
	year := context.Param("year")

	dateInt, err := strconv.Atoi(date)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	data, err := controller.hijriService.GregorianToHijri(dateInt, monthInt, yearInt)
	if err != nil {
		helper.BuildErrorResponse("Internal Server Error", err.Error(), nil)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}
