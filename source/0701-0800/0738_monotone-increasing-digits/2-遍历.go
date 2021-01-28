package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(156536329))
}

// leetcode738_单调递增的数字
func monotoneIncreasingDigits(N int) int {
	arr := []byte(strconv.Itoa(N))
	maxValue := -1
	index := -1
	for i := 0; i < len(arr)-1; i++ {
		if int(arr[i]-'0') > maxValue {
			maxValue = int(arr[i] - '0')
			index = i
		}
		if arr[i] > arr[i+1] {
			arr[index]--
			for j := index + 1; j < len(arr); j++ {
				arr[j] = '9'
			}
			break
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}
