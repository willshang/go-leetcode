package main

import (
	"fmt"
	"sort"
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
	tempArr := []byte(strconv.Itoa(num))
	sort.Slice(tempArr, func(i, j int) bool {
		return tempArr[i] > tempArr[j]
	})

	for i := 0; i < len(arr); i++ {
		if arr[i] != tempArr[i] {
			arr[i], arr[temp[int(tempArr[i]-'0')]] = arr[temp[int(tempArr[i]-'0')]], arr[i]
			res, _ = strconv.Atoi(string(arr))
			return res
		}
	}
	return res
}
