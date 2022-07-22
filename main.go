package main

import (
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

	r.Run(":3000")
}
