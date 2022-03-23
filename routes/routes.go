package routes

import (
	HelloWorldController "my_gym_go/controllers/hello_world"

	"github.com/labstack/echo/v4"
)

func InitHelloWorldRoutes(e *echo.Echo) {
	e.GET("hello", HelloWorldController.Hello)
	e.GET("world", HelloWorldController.World)
}
