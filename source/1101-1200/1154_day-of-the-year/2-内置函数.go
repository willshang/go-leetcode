package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}

func dayOfYear(date string) int {
	format := "2006-01-02"
	t, _ := time.Parse(format, date)
	return t.YearDay()
}
