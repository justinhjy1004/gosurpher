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

	var b []models.Blog = db.Select_blogs()

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
	return Render(c, views.BlogMainPage(b, summary))

}

// Handles generic blogs, with routes starting with generic
func GenericBlogHandler(c echo.Context) error {

	// Get the parameters of route and type
	route := c.Param("route")
	blog_type := c.Param("type")

	// Query blog by selecting the route
	var b models.Blog = db.Select_blog_by_route("generic/" + blog_type + "/" + route)

	// Split content into paragraphs
	var content []string = strings.Split(b.Content, "\n")
	var paragraphs []models.Paragraph
	// Parse the texts and links and populate the Paragraph object
	for i := 0; i < len(content); i++ {
		texts, links := TextLinkParser(content[i])
		paragraphs = append(paragraphs, models.Paragraph{Texts: texts, Urls: links})
	}

	return Render(c, views.BlogPage(b, paragraphs))

}
