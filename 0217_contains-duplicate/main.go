package main

import "fmt"

func main() {
	arr := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(arr))
}
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++{
		if _, ok := m[nums[i]]; ok{
			return true
		}else {
			m[nums[i]] = 1
		}
	}
	return false
}