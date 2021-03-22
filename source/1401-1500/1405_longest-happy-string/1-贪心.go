package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestDiverseString(1, 1, 7))
}

// leetcode1405_最长快乐字符串
func longestDiverseString(a int, b int, c int) string {
	arr := [][2]int{
		{0, a},
		{1, b},
		{2, c},
	}
	res := make([]byte, 0)
	for {
		// 按照次数排序
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][1] < arr[j][1]
		})
		// 每次放1个，如果最后2个相同，则使用次多的
		if len(res) >= 2 &&
			res[len(res)-1] == byte(arr[2][0]+'a') &&
			res[len(res)-2] == byte(arr[2][0]+'a') {
			if arr[1][1] > 0 { // 使用次多的
				res = append(res, byte(arr[1][0]+'a'))
				arr[1][1]--
			} else {
				break
			}
		} else {
			if arr[2][1] > 0 { // 使用最多的
				res = append(res, byte(arr[2][0]+'a'))
				arr[2][1]--
			} else {
				break
			}
		}
	}
	return string(res)
}
