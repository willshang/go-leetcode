package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	a, b := nums[0], nums[0]
	countA, countB := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == a {
			countA++
			continue
		}
		if nums[i] == b {
			countB++
			continue
		}
		if countA == 0 {
			a = nums[i]
			countA++
			continue
		}
		if countB == 0 {
			b = nums[i]
			countB++
			continue
		}
		countA--
		countB--
	}
	countA, countB = 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == a {
			countA++
		} else if nums[i] == b {
			countB++
		}
	}
	if countA > len(nums)/3 {
		res = append(res, a)
	}
	if countB > len(nums)/3 {
		res = append(res, b)
	}
	return res
}
