package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bestSeqAtIndex([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 0, 190, 95, 110}))
}

// 程序员面试金典17.08_马戏团人塔
func bestSeqAtIndex(height []int, weight []int) int {
	arr := make([][2]int, 0)
	for i := 0; i < len(height); i++ {
		arr = append(arr, [2]int{height[i], weight[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] > arr[j][0]
	})
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if len(res) == 0 || arr[res[len(res)-1]][0] > arr[i][0] &&
			arr[res[len(res)-1]][1] > arr[i][1] {
			res = append(res, i)
		} else {
			left := 0
			right := len(res) - 1
			for left <= right {
				mid := left + (right-left)/2
				if arr[res[mid]][0] > arr[i][0] && arr[res[mid]][1] > arr[i][1] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			res[left] = i
		}
	}
	return len(res)
}
