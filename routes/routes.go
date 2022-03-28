package routes

import (
	"my_gym_go/controllers"
	HelloWorldController "my_gym_go/controllers/hello_world"

	"github.com/labstack/echo/v4"
)

func InitHelloWorldRoutes(e *echo.Echo, controllers *controllers.ConfigurationControllers) {
	e.GET("hello", HelloWorldController.Hello)
	e.GET("world", HelloWorldController.World)
	e.GET("exercicios", controllers.ExercicioController.FindAll)
}
