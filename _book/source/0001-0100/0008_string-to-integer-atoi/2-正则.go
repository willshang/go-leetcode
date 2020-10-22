package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
}

func myAtoi(str string) int {
	re := regexp.MustCompile(`^[+-]?\d+`)
	arrS := re.FindAllString(strings.Trim(str, " "), -1)
	if len(arrS) == 0 {
		return 0
	}
	arr := arrS[0]
	res := 0
	isFlag := byte(' ')
	if !(arr[0] >= '0' && arr[0] <= '9') {
		isFlag = arr[0]
		arr = arr[1:]
	}
	for i := 0; i < len(arr); i++ {
		value := int(arr[i] - '0')
		if isFlag == '-' {
			if res > 214748364 || (res == 214748364 && value >= 8) {
				return math.MinInt32
			}
		} else if isFlag == ' ' || isFlag == '+' {
			if res > 214748364 || (res == 214748364 && value >= 7) {
				return math.MaxInt32
			}
		}
		res = res*10 + value
	}
	if isFlag == '-' {
		return -1 * res
	}
	return res
}
