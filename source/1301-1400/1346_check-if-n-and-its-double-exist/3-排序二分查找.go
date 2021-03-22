package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	//fmt.Println(checkIfExist([]int{-10, 12, -20, -8, 15}))
	fmt.Println(checkIfExist([]int{102, -592, 457, 802, 98, -132, 883, 356, -857, 461, -453, 522, 250, 476, 991, 540,
		-852, -485, -637, 999, -803, -691, -880, 881, -584, 750, -124, 745, -909, -892, 304, -814, 868, 665, 50, -40, 26, -242,
		-797, -360, -918, -741, 88, -933, -93, 360, -738, 833, -191, 563, 449, 840, 806, -87, -950, 508, 74, -448, -815, -488, 639, -334}))
}

func checkIfExist(arr []int) bool {
	var target int
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		left := i + 1
		right := len(arr) - 1
		if arr[i] >= 0 {
			target = 2 * arr[i]
		} else {
			if arr[i]%2 == -1 {
				continue
			}
			target = arr[i] / 2
		}
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] == target {
				return true
			} else if arr[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
