package models

type Date struct {
	year  int
	month int
	day   int
}

type blog_post struct {
	id      int
	title   string
	autor   string
	content string
	date    Date
	tags    []string
	status  string
}
