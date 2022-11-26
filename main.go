package main

import (
	"os"

	"github.com/fahmialfareza/dzikir-app-api/config"
	"github.com/fahmialfareza/dzikir-app-api/controller"
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/fahmialfareza/dzikir-app-api/repository"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	environment *entity.Config = config.SetupEnvironment()
	redisClient *redis.Client  = config.NewRedisConn(environment)

	salatTimeRepository repository.SalatTimeRepository = repository.NewSalatRepository(environment, redisClient)
	quranRepository     repository.QuranRepository     = repository.NewQuranRepository(environment, redisClient)
	hijriRepository     repository.HijriRepository     = repository.NewHijriRepository(environment, redisClient)
	geocodeRepository   repository.GeocodeRepository   = repository.NewGeocodeRepository(environment, redisClient)

	salatTimeService service.SalatTimeService = service.NewSalatTimeService(salatTimeRepository)
	quranService     service.QuranService     = service.NewQuranService(quranRepository)
	hijriService     service.HijriService     = service.NewHijriService(hijriRepository)
	geocodeService   service.GeocodeService   = service.NewGeocodeService(geocodeRepository)

	salatTimeController controller.SalatTimeController = controller.NewSalatTimeController(salatTimeService)
	quranController     controller.QuranController     = controller.NewQuranController(quranService)
	hijriController     controller.HijriController     = controller.NewHijriController(hijriService)
	geocodeController   controller.GeocodeController   = controller.NewGeocodeController(geocodeService)
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

	hijriRoutes := r.Group("/api/v1/hijri")
	{
		hijriRoutes.GET("/:date/:month/:year", hijriController.GregorianToHijri)
	}

	geocodeRoutes := r.Group("/api/v1/geocode")
	{
		geocodeRoutes.GET("/reverse/:lat/:lng", geocodeController.ReverseGeocode)
	}

	r.NoRoute(func(context *gin.Context) {
		res := helper.BuildErrorResponse("Not Found", "Resource not found", nil)
		context.JSON(404, res)
	})

	r.NoMethod(func(context *gin.Context) {
		res := helper.BuildErrorResponse("Method Not Allowed", "Method not allowed", nil)
		context.JSON(405, res)
	})

	if os.Getenv("PORT") != "" {
		// Heroku add a env variable called PORT, if exist we will use it
		r.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		// If is running on localhost (our computer), no PORT env variable
		r.Run("0.0.0.0:3000")
	}
}
