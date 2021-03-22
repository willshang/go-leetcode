package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
}

func daysBetweenDates(date1 string, date2 string) int {
	v1 := totalDay(date1)
	v2 := totalDay(date2)
	if v1 > v2 {
		return v1 - v2
	}
	return v2 - v1
}

var monthDate = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func totalDay(date string) int {
	var year, month, day int
	arr := strings.Split(date, "-")
	year, _ = strconv.Atoi(arr[0])
	month, _ = strconv.Atoi(arr[1])
	day, _ = strconv.Atoi(arr[2])
	total := 0
	for i := 1971; i < year; i++ {
		total = total + 365
		if isLeap(i) {
			total = total + 1
		}
	}
	for i := 0; i < month-1; i++ {
		total = total + monthDate[i]
		if i == 1 && isLeap(year) {
			total = total + 1
		}
	}
	total = total + day
	return total
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
