package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type QuranController interface {
	AllChapters(ctx *gin.Context)
	AllVerses(ctx *gin.Context)
	VersesByChapter(ctx *gin.Context)
}

type quranController struct {
	quranService service.QuranService
}

func NewQuranController(quranService service.QuranService) QuranController {
	return &quranController{
		quranService: quranService,
	}
}

func (controller *quranController) AllChapters(ctx *gin.Context) {
	data, err := controller.quranService.AllChapters(ctx)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}

func (controller *quranController) VersesByChapter(ctx *gin.Context) {
	chapter := ctx.Param("chapter")

	chapterInt, err := strconv.Atoi(chapter)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Chapter is required", nil)
		ctx.JSON(400, res)
		return
	}

	data, err := controller.quranService.VersesByChapter(ctx, chapterInt)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}

func (controller *quranController) AllVerses(ctx *gin.Context) {
	data, err := controller.quranService.AllVerses(ctx)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		ctx.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	ctx.JSON(200, res)
}
