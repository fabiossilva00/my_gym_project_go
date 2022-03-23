package main

import (
	"log"
	"my_gym_go/db"
	"my_gym_go/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitHelloWorldRoutes(e)

	_, err := db.InitMongoDb()

	if err != nil {
		log.Panicln(err)
	}

	e.Logger.Fatal(e.Start(":1323"))

}
