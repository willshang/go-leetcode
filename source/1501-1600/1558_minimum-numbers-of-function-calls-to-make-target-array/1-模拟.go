package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{2, 2}))
}

func minOperations(nums []int) int {
	res := 0
	for judge(nums) != true {
		res++ // 最后一次循环会多算一次,最后要减掉
		res = res + div(nums)
	}
	return res - 1
}

func judge(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func div(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		if arr[i]%2 == 1 {
			res++
		}
		arr[i] = arr[i] / 2
	}
	return res
}
