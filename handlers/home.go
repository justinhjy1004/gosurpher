package handlers

import (
	"gosurpher/views"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {

	return Render(c, views.Home("Gosurpher | Home"))

}
