package main

import (
	"gosurpher/handlers"
	"gosurpher/models"
	"gosurpher/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func blogHandler(c echo.Context) error {

	b := models.Blog{
		Id:      1,
		Autor:   "justino",
		Title:   "test",
		Content: "go",
	}

	return render(c, views.Blog(b))

}

func render(c echo.Context, cmp templ.Component) error {
	// This is real
	return cmp.Render(c.Request().Context(), c.Response())

}

func main() {

	e := echo.New()

	// Static CSS
	e.Static("/", "static/")

	e.GET("/", handlers.HomeHandler)
	e.GET("/blog", blogHandler)
	e.Logger.Fatal(e.Start(":42069"))

}
