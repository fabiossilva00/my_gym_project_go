package HelloWorldController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func World(c echo.Context) error {
	return c.String(http.StatusOK, " World!")
}
