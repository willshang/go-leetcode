package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)

}

func sortColors(nums []int) {
	arr := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]]++
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			nums[count] = i
			count++
		}
	}
}
