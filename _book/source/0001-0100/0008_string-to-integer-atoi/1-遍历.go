package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))
}

func myAtoi(str string) int {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	str = str[i:]
	arr := make([]byte, 0)
	isFlag := byte(' ')
	for j := 0; j < len(str); j++ {
		if str[j] >= '0' && str[j] <= '9' {
			arr = append(arr, str[j])
		} else {
			if len(arr) > 0 {
				break
			}
			if str[j] != ' ' && str[j] != '+' && str[j] != '-' {
				return 0
			}
			if isFlag != ' ' {
				return 0
			}
			isFlag = str[j]
		}
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		value := int(arr[i] - '0')
		res = res*10 + value
		if isFlag == '-' {
			if -1*res < math.MinInt32 {
				return math.MinInt32
			}
		} else if isFlag == ' ' || isFlag == '+' {
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}
	if isFlag == '-' {
		return -1 * res
	}
	return res
}
