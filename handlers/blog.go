package handlers

import (
	"gosurpher/models"
	"gosurpher/views"
	"time"

	"github.com/labstack/echo/v4"
)

func BlogHandler(c echo.Context) error {

	b1 := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Autor:   "justino",
		Title:   "test",
		Content: "go",
	}

	b2 := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Autor:   "justino",
		Title:   "test",
		Content: "go",
	}

	b3 := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Autor:   "justino",
		Title:   "test",
		Content: "go",
	}

	b := []models.Blog{b1, b2, b3}

	return Render(c, views.Blog(b))

}
