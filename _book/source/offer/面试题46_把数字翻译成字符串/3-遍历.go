package main

import (
	"fmt"
)

func main() {
	fmt.Println(translateNum(10022))
}

func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	arr := make([]int, 1)
	arr[0] = 1
	i := 0
	prev := -1
	for num > 0 {
		i++
		arr = append(arr, 0)
		arr[i] = arr[i-1]
		digit1 := num % 10
		num = num / 10
		if digit1 != 0 && prev >= 0 && digit1*10+prev <= 25 {
			arr[i] = arr[i] + arr[i-2]
		}
		prev = digit1
	}
	return arr[i]
}
