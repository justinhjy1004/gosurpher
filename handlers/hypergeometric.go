package handlers

import (
	"gosurpher/db"
	"gosurpher/models"
	"gosurpher/views"

	"github.com/labstack/echo/v4"
)

func HypergeometricHandler(c echo.Context) error {

	var b models.Blog = db.Select_blog_by_route("hypergeometric")

	return Render(c, views.Hypergeometric(b))

}
