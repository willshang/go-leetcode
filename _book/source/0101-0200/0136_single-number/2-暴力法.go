package main

import "fmt"

func main() {
	//arr := []int{4, 1, 2, 1, 2}
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				flag = true
				break
			}
		}
		if flag == false {
			return nums[i]
		}

	}
	return -1
}
