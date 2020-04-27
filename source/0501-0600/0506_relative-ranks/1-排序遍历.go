package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	nums := []int{10, 100, 5, 2, 9}
	//nums := []int{5,4,3,2,1}
	fmt.Println(findRelativeRanks(nums))
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}

// leetcode506_相对名次
func findRelativeRanks(nums []int) []string {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	m := make(map[int]string)
	for i := 0; i < len(temp); i++ {
		if i == len(temp)-1 {
			m[temp[i]] = "Gold Medal"
		} else if i == len(temp)-2 {
			m[temp[i]] = "Silver Medal"
		} else if i == len(temp)-3 {
			m[temp[i]] = "Bronze Medal"
		} else {
			m[temp[i]] = strconv.Itoa(len(temp) - i)
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, m[nums[i]])
	}
	return res
}
