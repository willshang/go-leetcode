package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(29, 2, 2016))
}

// leetcode1185_一周中的第几天
func dayOfTheWeek(day int, month int, year int) string {
	t, _ := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", year, month, day))
	return t.Weekday().String()
}
