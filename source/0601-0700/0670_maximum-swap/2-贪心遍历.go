package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(2736))
}

func maximumSwap(num int) int {
	if num <= 11 {
		return num
	}
	res := num
	arr := []byte(strconv.Itoa(num))
	temp := [10]int{}
	for i := 0; i < len(arr); i++ {
		temp[arr[i]-'0'] = i // 每个数字最后一次出现的位置
	}
	for i := 0; i < len(arr); i++ {
		// 寻找最后面比当前数字大并且最大数字并进行交换
		for j := 9; j > int(arr[i]-'0'); j-- {
			if temp[j] > i {
				arr[i], arr[temp[j]] = arr[temp[j]], arr[i]
				res, _ = strconv.Atoi(string(arr))
				return res
			}
		}
	}
	return res
}
