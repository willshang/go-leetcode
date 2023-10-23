package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(156536329))
}

func monotoneIncreasingDigits(N int) int {
	arr := []byte(strconv.Itoa(N))
	i := 1
	for i < len(arr) && arr[i-1] <= arr[i] {
		i++
	}
	// 前面有逆序的
	if i < len(arr) {
		// 前面减去1, 如：332=>2xx要减2次
		for i > 0 && arr[i] < arr[i-1] {
			arr[i-1]--
			i--
		}
		i++
		for ; i < len(arr); i++ {
			arr[i] = '9'
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}
