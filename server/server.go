package server

import (
	"my_gym_go/config"
	"my_gym_go/controllers"
	"my_gym_go/db"
	"my_gym_go/model"
	"my_gym_go/repository"
	"my_gym_go/routes"
	"my_gym_go/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	configuration *config.Configuration
	clientMongo   *mongo.Client
)

func init() {
	configuration = config.LoadConfig()
	clientMongo = db.InitMongoDb(configuration)
}

func NewServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	valid := validator.New()
	e.Validator = model.NewCustomValidator(valid)

	database := db.GetDatabase(configuration.NameDatabase)
	repos := repository.NewConfigurationRepositories(database)
	service := service.NewConfigurationService(repos)
	controllers := controllers.NewConfigurationControllers(service)
	routes.InitHelloWorldRoutes(e, controllers)

	e.Logger.Fatal(e.Start(":1323"))

}
