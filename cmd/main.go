package main

import (
	"gosurpher/handlers"

	"github.com/labstack/echo/v4"
)

/* This is the routing map from a given GET endpoint
call to the function that renders it. Note that
unique pages have their own handlers, there are
generic handlers that handle repetitive pages (style-wise) */

var route = map[string]interface{}{
	"/":                     handlers.HomeHandler,
	"/blog":                 handlers.BlogHandler,
	"/hypergeometric":       handlers.HypergeometricHandler,
	"/generic/:type/:route": handlers.GenericBlogHandler,
}

func main() {

	e := echo.New()

	// Static CSS
	// Note that assets and css files live here
	e.Static("/", "static/")

	// Rendering all pages per the routes to each page
	for r, h := range route {
		e.GET(r, h.(func(c echo.Context) error))
	}

	e.Logger.Fatal(e.Start(":42069"))

}
