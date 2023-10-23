package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findRadius([]int{4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8}))
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
}

func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 {
		return 0
	}
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	length := len(heaters)
	for i := 0; i < len(houses); i++ {
		left := 0
		right := length - 1
		for left < right {
			mid := left + (right-left)/2
			if heaters[mid] < houses[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		dis := 0
		if heaters[left] < houses[i] {
			dis = houses[i] - heaters[left]
		} else if heaters[left] > houses[i] {
			if left == 0 {
				dis = heaters[0] - houses[i]
			} else {
				dis = Min(heaters[left]-houses[i], houses[i]-heaters[left-1])
			}
		}
		res = Max(res, dis)
	}
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
