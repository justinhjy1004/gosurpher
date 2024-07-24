package handlers

import (
	"gosurpher/models"
	"gosurpher/views"
	"time"

	"github.com/labstack/echo/v4"
)

func HypergeometricHandler(c echo.Context) error {

	b := models.Blog{
		Id:      1,
		Date:    time.Date(2024, time.July, 24, 0, 0, 0, 0, time.UTC),
		Autor:   "Justin Ho",
		Title:   "The Hypergeometric Story",
		Content: "Inspired by the favourite statistics teacher Blitzstein, every probability distribution has its own story, and this is the first of my series in telling the hypergeometric distribution. One that I have recently learned, but one that I have a fun story to tell.",
		Route:   "hypergeometric",
	}

	return Render(c, views.Hypergeometric(b))

}
