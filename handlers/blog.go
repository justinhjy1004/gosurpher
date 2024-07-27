package handlers

import (
	"gosurpher/db"
	"gosurpher/views"

	"github.com/labstack/echo/v4"
)

func BlogHandler(c echo.Context) error {

	b := db.Select_blogs()

	return Render(c, views.Blog(b))

}
