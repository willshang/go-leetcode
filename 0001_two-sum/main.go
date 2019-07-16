package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

//执行用时：4 ms
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		if num, ok := m[b]; ok && num != i {
			return []int{i, m[b]}
		}
	}
	return nil
}

//执行用时：120 ms
/*func twoSum(nums []int, target int) []int  {
	length := len(nums)
	for i := 0; i < length; i++{
		for j := i + 1; j < length; j++{
			if nums[j] == target - nums[i]{
				return []int{i,j}
			}
		}
	}
	return nil
}*/

//执行用时：4 ms
/*func twoSum(nums []int, target int) []int  {
	m := make(map[int]int,len(nums))
	for i, b := range nums{
		if j, ok := m[target-b]; ok{
			return []int{j,i}
		}
		m[b] = i
	}
	return nil
}*/
