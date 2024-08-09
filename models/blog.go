package models

import (
	"time"
)

/* Basic Blog structure */
type Blog struct {
	Id      int
	Date    time.Time
	Title   string
	Autor   string
	Content string
	Route   string
	Type    string
}

type Link struct {
	Text string
	Url  string
}

/*
Paragraph structure to render in Generic Blogs
*/
type Paragraph struct {
	Texts []string
	Urls  []Link
}
