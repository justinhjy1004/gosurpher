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
		Date:    time.Date(2024, time.July, 24, 0, 0, 0, 0, time.UTC),
		Autor:   "Justin Ho",
		Title:   "The Hypergeometric Story",
		Content: "Inspired by the favourite statistics teacher Blitzstein, every probability distribution has its own story, and this is the first of my series in telling the hypergeometric distribution. One that I have recently learned, but one that I have a fun story to tell.",
		Route:   "hypergeometric",
	}

	b2 := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Autor:   "justino",
		Title:   "test",
		Content: "go",
		Route:   "test",
	}

	b3 := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Autor:   "justino",
		Title:   "test",
		Content: "go",
		Route:   "test",
	}

	b := []models.Blog{b1, b2, b3}

	return Render(c, views.Blog(b))

}
