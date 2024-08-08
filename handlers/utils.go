package handlers

import (
	"gosurpher/models"
	"regexp"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, cmp templ.Component) error {

	return cmp.Render(c.Request().Context(), c.Response())

}

// This is mainly to center to go-echarts when rendered
func RenderGoEcharts(chart render.ChartSnippet) string {
	var element string = chart.Element
	var script string = chart.Script

	element = strings.Replace(element, "width:900px;height:500px;", "width:900px;height:500px;margin-left:auto;margin-right:auto;", 1)
	echart_html := element + script

	return (echart_html)

}

func TextLinkParser(txt string) ([]string, []models.Link) {
	// Regex for link recognition
	re := regexp.MustCompile(`\(\s*([^,]+?)\s*,\s*(http[s]?:\/\/[^\s]+)\s*\)`)

	// Split the text by the links
	split := re.Split(txt, -1)

	urls := re.FindAllString(txt, -1)

	links := make([]models.Link, 0)

	for i := range urls {
		url := re.FindStringSubmatch(urls[i])
		link := models.Link{Text: url[1], Url: url[2]}
		links = append(links, link)
	}

	return split, links

}
