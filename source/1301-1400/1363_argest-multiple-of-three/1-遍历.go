package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestMultipleOfThree([]int{8, 1, 9}))
	fmt.Println(largestMultipleOfThree([]int{1}))
}

// leetcode1363_形成三的最大倍数
func largestMultipleOfThree(digits []int) string {
	arr, arr3 := [10]int{}, [3]int{}
	var sum, index, count int // 需要移除的数字和个数
	for i := 0; i < len(digits); i++ {
		arr[digits[i]]++
		arr3[digits[i]%3]++
		sum = sum + digits[i]
	}
	if sum%3 == 1 { // 多个1
		if arr3[1] >= 1 {
			index = 1
			count = 1
		} else { // 移除2个2
			index = 2
			count = 2
		}
	} else if sum%3 == 2 {
		if arr3[2] >= 1 {
			index = 2
			count = 1
		} else { // 移除2个1
			index = 1
			count = 2
		}
	}
	res := make([]byte, 0)
	for i := 0; i <= 9; i++ {
		for j := 0; j < arr[i]; j++ {
			if count > 0 && i%3 == index {
				count--
			} else {
				res = append(res, byte('0'+i))
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	if len(res) > 0 && res[0] == '0' {
		return "0"
	}
	return string(res)
}
