package main

import "fmt"

func main() {
	fmt.Println(getTriggerTime([][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}, [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}))
}

func getTriggerTime(increase [][]int, requirements [][]int) []int {
	for i := 1; i < len(increase); i++ {
		increase[i][0] = increase[i][0] + increase[i-1][0]
		increase[i][1] = increase[i][1] + increase[i-1][1]
		increase[i][2] = increase[i][2] + increase[i-1][2]
	}
	res := make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		C, R, H := requirements[i][0], requirements[i][1], requirements[i][2]
		if C == 0 && R == 0 && H == 0 {
			res[i] = 0
			continue
		}
		if C > increase[len(increase)-1][0] ||
			R > increase[len(increase)-1][1] ||
			H > increase[len(increase)-1][2] {
			res[i] = -1
			continue
		}
		left, right := 0, len(increase)-1
		index := -1
		for left <= right {
			mid := left + (right-left)/2
			if increase[mid][0] >= C && increase[mid][1] >= R && increase[mid][2] >= H {
				index = mid + 1
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		res[i] = index
	}
	return res
}
