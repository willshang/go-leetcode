package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{
		{1, 4, 1, 40, 10},
		{2, 8, 0, 50, 5},
		{3, 8, 1, 30, 4},
		{4, 10, 0, 10, 3},
		{5, 1, 1, 15, 1},
	}
	fmt.Println(filterRestaurants(arr, 1, 50, 10))
}

// leetcode1333_餐厅过滤器
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	res := make([]int, 0)
	arr := make([][2]int, 0)
	for i := 0; i < len(restaurants); i++ {
		if restaurants[i][3] > maxPrice || restaurants[i][4] > maxDistance {
			continue
		}
		if veganFriendly == 1 {
			if restaurants[i][2] == 0 {
				continue
			}
		}
		arr = append(arr, [2]int{restaurants[i][0], restaurants[i][1]})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] > arr[j][0]
		}
		return arr[i][1] > arr[j][1]
	})
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[i][0])
	}
	return res
}
