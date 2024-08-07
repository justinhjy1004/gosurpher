package handlers

import (
	"gosurpher/db"
	"gosurpher/models"
	"gosurpher/views"
	"strings"

	"github.com/labstack/echo/v4"
)

// This Renders the 'front page' of the blog
func BlogHandler(c echo.Context) error {

	var b []models.Blog

	b = db.Select_blogs()

	// Create a 300 character long summary for rendering in the card
	var summary []string
	var max_length int = 300

	for i := range b {
		if len(b[i].Content) < max_length {
			summary = append(summary, b[i].Content)
		} else {
			summary = append(summary, b[i].Content[:max_length]+"...")
		}
	}
	return Render(c, views.Blog(b, summary))

}

// Handles generic blogs, with routes starting with generic
func GenericBlogHandler(c echo.Context) error {

	route := c.Param("route")
	blog_type := c.Param("type")

	var b models.Blog

	b = db.Select_blog_by_route("generic/" + blog_type + "/" + route)

	var content []string
	content = strings.Split(b.Content, "\n")

	return Render(c, views.BlogPage(b, content))

}
