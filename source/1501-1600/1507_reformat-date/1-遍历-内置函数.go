package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
}

// leetcode1507_转变日期格式
func reformatDate(date string) string {
	month := []string{"", "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	arr := strings.Split(date, " ")
	res := arr[2] + "-"
	for i := 1; i < len(month); i++ {
		if month[i] == arr[1] {
			res = res + fmt.Sprintf("%02d", i) + "-"
		}
	}
	/*
		if arr[0][1] >= '0' && arr[0][1] <= '9' {
			res = res + arr[0][:2]
		} else {
			res = res + "0" + arr[0][:1]
		}
	*/
	if len(arr[0]) == 3 {
		res = res + "0" + arr[0][:1]
	} else {
		res = res + arr[0][:2]
	}
	return res
}
