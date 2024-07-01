package handlers

import (
	"gosurpher/models"
	"gosurpher/views"

	"github.com/labstack/echo/v4"
)

func BlogHandler(c echo.Context) error {

	b := models.Blog{
		Id:      1,
		Autor:   "justino",
		Title:   "test",
		Content: "go",
	}

	return Render(c, views.Blog(b))

}
