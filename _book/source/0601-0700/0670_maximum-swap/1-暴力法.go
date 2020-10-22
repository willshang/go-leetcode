package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(2736))
}

// leetcode670_最大交换
func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	res := num
	arr := []byte(strconv.Itoa(num))
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			tempArr := make([]byte, len(arr))
			copy(tempArr, arr)
			tempArr[i], tempArr[j] = tempArr[j], tempArr[i]
			newValue, _ := strconv.Atoi(string(tempArr))
			if newValue > res {
				res = newValue
			}
		}
	}
	return res
}
