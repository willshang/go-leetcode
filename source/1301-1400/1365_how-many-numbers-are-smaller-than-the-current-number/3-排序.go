package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	m := make(map[int]int)
	count := 0
	m[temp[0]] = count

	for i := 1; i < len(temp); i++ {
		count++
		if temp[i-1] != temp[i] {
			m[temp[i]] = count
		} else {
			m[temp[i]] = m[temp[i-1]]
		}
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = m[nums[i]]
	}
	return res
}
