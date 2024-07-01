package main

import (
	"gosurpher/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Static CSS
	e.Static("/", "static/")

	e.GET("/", handlers.HomeHandler)
	e.GET("/blog", handlers.BlogHandler)
	e.Logger.Fatal(e.Start(":42069"))

}
