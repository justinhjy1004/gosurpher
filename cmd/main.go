package main

import (
	"gosurpher/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {

	return render(c, views.Home("Page Title"))

}

func blogHandler(c echo.Context) error {

	return render(c, views.Blog())

}

func render(c echo.Context, cmp templ.Component) error {
	// This is real
	return cmp.Render(c.Request().Context(), c.Response())

}

func main() {

	e := echo.New()
	e.GET("/", homeHandler)
	e.GET("/blog", blogHandler)
	e.Logger.Fatal(e.Start(":42069"))

}
