package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

var m map[int]int

func combinationSum4(nums []int, target int) int {
	m = make(map[int]int)
	res := dfs(nums, target)
	if res == -1 {
		return 0
	}
	return res
}

func dfs(nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return -1
	}
	if v, ok := m[target]; ok {
		return v
	}
	temp := 0
	for i := 0; i < len(nums); i++ {
		if dfs(nums, target-nums[i]) != -1 {
			temp = temp + dfs(nums, target-nums[i])
		}
	}
	m[target] = temp
	return temp
}
