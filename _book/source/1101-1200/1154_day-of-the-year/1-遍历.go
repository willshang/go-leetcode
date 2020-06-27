package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}

// leetcode1154_一年中的第几天
func dayOfYear(date string) int {
	arr := strings.Split(date, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])
	res := 0
	for i := 0; i < month; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			res = res + 31
		case 4, 6, 9, 11:
			res = res + 30
		case 2:
			res = res + 28
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				res = res + 1
			}
		}
	}
	res = res + day
	return res
}
