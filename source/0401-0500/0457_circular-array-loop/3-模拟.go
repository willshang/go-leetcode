package main

import "fmt"

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
}

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		if judge(nums, n, i) == true {
			return true
		}
	}
	return false
}

func judge(nums []int, n, cur int) bool {
	start := cur
	dir := 1
	if nums[cur] < 0 {
		dir = -1
	}
	count := 1
	for {
		if count > n {
			return false
		}
		next := ((cur+nums[cur])%n + n) % n
		if (dir > 0 && nums[next] < 0) || (dir < 0 && nums[next] > 0) {
			return false
		}
		if next == start { // 走到起点
			return count > 1
		}
		count++
		cur = next
	}
}
