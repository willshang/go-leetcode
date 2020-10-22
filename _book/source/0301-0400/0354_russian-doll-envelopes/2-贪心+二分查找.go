package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	// 宽[0] 高[1]排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	arr := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		left := 0
		right := len(arr) - 1
		for left <= right {
			mid := left + (right-left)/2
			if envelopes[i][1] > arr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left >= len(arr) {
			arr = append(arr, envelopes[i][1])
		} else {
			arr[left] = envelopes[i][1]
		}
	}
	return len(arr)
}
