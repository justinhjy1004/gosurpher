package main

import (
	"gosurpher/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Static CSS
	e.Static("/", "static/")

	// Rendering all pages per the routes to each page
	for r, h := range route {
		e.GET(r, h.(func(c echo.Context) error))
	}

	e.Logger.Fatal(e.Start(":42069"))

}

var route = map[string]interface{}{
	"/":               handlers.HomeHandler,
	"/blog":           handlers.BlogHandler,
	"/hypergeometric": handlers.HypergeometricHandler,
	":type/:route":    handlers.GenericBlogHandler,
}
