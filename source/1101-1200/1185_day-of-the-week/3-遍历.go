package main

import (
	"fmt"
)

func main() {
	//fmt.Println(dayOfTheWeek(31, 8, 2019))
	//fmt.Println(dayOfTheWeek(18, 7, 1999))
	//fmt.Println(dayOfTheWeek(29, 2, 2016))
	//fmt.Println(dayOfTheWeek(16, 8, 1993))
	//fmt.Println(dayOfTheWeek(15, 8, 1993))
	//fmt.Println(dayOfTheWeek(14, 8, 1993))
	//fmt.Println(dayOfTheWeek(1, 1, 1971))
	fmt.Println(dayOfTheWeek(1, 1, 1972))
	//fmt.Println(dayOfTheWeek(1, 1, 1973))
	//fmt.Println(dayOfTheWeek(1, 1, 1974))
}

var arr = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDate = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func dayOfTheWeek(day int, month int, year int) string {
	day1 := totalDay(1993, 8, 15)
	day2 := totalDay(year, month, day)
	diff := 6 - day1%7
	return arr[(day2+diff)%7]
}

func totalDay(year, month, day int) int {
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
