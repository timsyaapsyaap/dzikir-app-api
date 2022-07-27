package main

import (
	"os"

	"github.com/fahmialfareza/dzikir-app-api/config"
	"github.com/fahmialfareza/dzikir-app-api/controller"
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var (
	environment         *entity.Config                 = config.SetupEnvironment()
	pool                *redis.Pool                    = config.NewPool(*environment)
	salatTimeRepository repository.SalatTimeRepository = repository.NewSalatRepository(environment)
	quranRepository     repository.QuranRepository     = repository.NewQuranRepository(environment, pool)
	salatTimeService    service.SalatTimeService       = service.NewSalatTimeService(salatTimeRepository)
	quranService        service.QuranService           = service.NewQuranService(quranRepository, pool)
	salatTimeController controller.SalatTimeController = controller.NewSalatTimeController(salatTimeService)
	quranController     controller.QuranController     = controller.NewQuranController(quranService)
)

func main() {
	r := gin.Default()

	salatTimeRoutes := r.Group("/api/v1/salat-time")
	{
		salatTimeRoutes.GET("/cities", salatTimeController.AllCities)
		salatTimeRoutes.GET("/cities/:id", salatTimeController.CityDetails)
		salatTimeRoutes.GET("/cities/find/:city", salatTimeController.FindCity)
		salatTimeRoutes.GET("/schedule/:cityId/:year/:month/:date", salatTimeController.Schedule)
	}

	quranRoutes := r.Group("/api/v1/quran")
	{
		quranRoutes.GET("/chapters", quranController.AllChapters)
		quranRoutes.GET("/verses", quranController.AllVerses)
		quranRoutes.GET("/verses/:chapter", quranController.VersesByChapter)
	}

	if os.Getenv("PORT") != "" {
		// Heroku add a env variable called PORT, if exist we will use it
		r.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		// If is running on localhost (our computer), no PORT env variable
		r.Run("0.0.0.0:3000")
	}
}
