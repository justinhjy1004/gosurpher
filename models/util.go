package models

import (
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d Date) as_string() string {
	return fmt.Sprintln("%d-%d-%d", d.year, d.month, d.day)
}
