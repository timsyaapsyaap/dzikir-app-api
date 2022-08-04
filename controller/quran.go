package controller

import (
	"strconv"

	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

type QuranController interface {
	AllChapters(context *gin.Context)
	AllVerses(context *gin.Context)
	VersesByChapter(context *gin.Context)
}

type quranController struct {
	quranService service.QuranService
}

func NewQuranController(quranService service.QuranService) QuranController {
	return &quranController{
		quranService: quranService,
	}
}

func (controller *quranController) AllChapters(context *gin.Context) {
	data, err := controller.quranService.AllChapters()
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *quranController) VersesByChapter(context *gin.Context) {
	chapter := context.Param("chapter")

	chapterInt, err := strconv.Atoi(chapter)
	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", "Chapter is required", nil)
		context.JSON(400, res)
		return
	}

	data, err := controller.quranService.VersesByChapter(chapterInt)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}

func (controller *quranController) AllVerses(context *gin.Context) {
	data, err := controller.quranService.AllVerses()
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", "Terjadi Kesalahan", nil)
		context.JSON(500, res)
		return
	}

	res := helper.BuildResponse(true, "OK", data)
	context.JSON(200, res)
}
