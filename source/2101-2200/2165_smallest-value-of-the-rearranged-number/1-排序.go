package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(smallestNumber(310))
	fmt.Println(smallestNumber(0))
}

// leetcode2165_重排数字的最小值
func smallestNumber(num int64) int64 {
	arr := []byte(strconv.FormatInt(num, 10))
	flag := int64(1)
	if num <= 0 {
		flag = -1
		arr = arr[1:]
	}
	if flag == 1 { // 正数
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		i := 0
		for arr[i] == '0' {
			i++
		}
		arr[0], arr[i] = arr[i], arr[0] // 交换第一个非0
	} else {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	}
	res, _ := strconv.ParseInt(string(arr), 10, 64)
	return flag * res
}
