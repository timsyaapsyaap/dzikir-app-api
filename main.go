package main

import (
	"os"

	"github.com/fahmialfareza/dzikir-app-api/config"
	"github.com/fahmialfareza/dzikir-app-api/controller"
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
	"github.com/fahmialfareza/dzikir-app-api/service"
	"github.com/gin-gonic/gin"
)

var (
	environment         *entity.Config                 = config.SetupEnvironment()
	salatTimeRepository repository.SalatTimeRepository = repository.NewSalatRepository(environment)
	salatTimeService    service.SalatTimeService       = service.NewSalatTimeService(salatTimeRepository)
	salatTimeController controller.SalatTimeController = controller.NewSalatTimeController(salatTimeService)
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

	if os.Getenv("PORT") != "" {
		// Heroku add a env variable called PORT, if exist we will use it
		r.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		// If is running on localhost (our computer), no PORT env variable
		r.Run("0.0.0.0:3000")
	}
}
