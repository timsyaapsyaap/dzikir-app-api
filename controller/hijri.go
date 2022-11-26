package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type HijriController interface {
	GregorianToHijri(ctx *gin.Context)
}

type hijriController struct {
	hijriService service.HijriService
}

func NewHijriController(hijriService service.HijriService) HijriController {
	return &hijriController{
		hijriService: hijriService,
	}
}

func (controller *hijriController) GregorianToHijri(ctx *gin.Context) {
	date := ctx.Param("date")
	month := ctx.Param("month")
	year := ctx.Param("year")

	dateInt, err := strconv.Atoi(date)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Date is required", nil)
		ctx.JSON(400, res)
		return
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Month is required", nil)
		ctx.JSON(400, res)
		return
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Year is required", nil)
		ctx.JSON(400, res)
		return
	}

	data, err := controller.hijriService.GregorianToHijri(ctx, dateInt, monthInt, yearInt)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}
