package models

import (
	"time"
)

type Blog struct {
	Id      int
	Date    time.Time
	Title   string
	Autor   string
	Content string
	Route   string
	Type    string
}
