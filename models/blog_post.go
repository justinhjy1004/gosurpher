package models

type Date struct {
	year  int
	month int
	day   int
}

type Blog struct {
	Id      int
	Title   string
	Autor   string
	Content string
}
