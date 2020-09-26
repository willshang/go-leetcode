package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

// leetcode1323_6和9组成的最大数字
func maximum69Number(num int) int {
	arr := []byte(strconv.Itoa(num))
	for i := 0; i < len(arr); i++ {
		if arr[i] == '6' {
			arr[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}
