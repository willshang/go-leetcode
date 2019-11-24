package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 2}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	newNums := make([]int, len(nums))
	var repeatNum int
	for _, v := range nums {
		if newNums[v-1] != 0 {
			repeatNum = v
		}
		newNums[v-1] = v
	}

	for i, v := range newNums {
		if v == 0 {
			return []int{repeatNum, i + 1}
		}
	}

	return []int{0, 0}
}

/*func findErrorNums(nums []int) []int {
	dup := 0
	for i := 0; i < len(nums); i++{
		n := abs(nums[i])
		if nums[n-1] < 0{
			dup = n
		}else {
			nums[n-1] = -nums[n-1]
		}
	}

	mis := 0
	for i, v := range nums{
		if v > 0{
			mis = i + 1
			break
		}
	}
	return []int{dup,mis}
}

func abs(a int) int  {
	if a > 0{
		return a
	}
	return -a
}*/
