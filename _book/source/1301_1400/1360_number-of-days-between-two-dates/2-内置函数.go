package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
}

// leetcode1360_日期之间隔几天
func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)
	value := int(t1.Sub(t2).Hours() / 24)
	if value > 0 {
		return value
	}
	return -value
}
