package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(10022))
}

func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	return getTranslateNum(strconv.Itoa(num))
}

func getTranslateNum(str string) int {
	length := len(str)
	arr := make([]int, length)
	count := 0
	for i := length - 1; i >= 0; i-- {
		count = 0
		if i < length-1 {
			count = arr[i+1]
		} else {
			count = 1
		}
		if i < length-1 {
			digit1 := str[i] - '0'
			digit2 := str[i+1] - '0'
			value := digit1*10 + digit2
			if value >= 10 && value <= 25 {
				if i < length-2 {
					count += arr[i+2]
				} else {
					count += 1
				}
			}
		}
		arr[i] = count
	}
	return arr[0]
}
