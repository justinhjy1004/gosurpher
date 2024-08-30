package main

import (
	"gosurpher/handlers"
	"os"

	"github.com/labstack/echo/v4"
)

/* This is the routing map from a given GET endpoint
call to the function that renders it. Note that
unique pages have their own handlers, there are
generic handlers that handle repetitive pages (style-wise) */

var route = map[string]interface{}{
	"/blog":                 handlers.HomeHandler,
	"/":                     handlers.BlogHandler,
	"/generic/:type/:route": handlers.GenericBlogHandler,
	"poisson":               handlers.PoissonHandler,
	"gamma":                 handlers.GammaHandler,
	"exponential":           handlers.ExponentialHandler,
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "42069"
	}

	e.Logger.Fatal(e.Start("0.0.0.0:" + port))

}
