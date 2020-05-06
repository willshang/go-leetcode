package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 2}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	n := len(nums)
	sum := 0
	repeatNum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := m[nums[i]]; ok {
			repeatNum = nums[i]
		}
		m[nums[i]] = 1
	}
	return []int{repeatNum, n*(n+1)/2 - sum + repeatNum}
}
