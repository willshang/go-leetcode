package main

import "fmt"

func main() {
	fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5}))
}

// 程序员面试金典08.03_魔术索引
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		}
	}
	return -1
}
