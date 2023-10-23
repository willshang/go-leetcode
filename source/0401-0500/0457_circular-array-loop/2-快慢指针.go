package main

import "fmt"

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
}

// leetcode457_环形数组循环
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		slow, fast := i, getNext(nums, n, i)
		// 保证是同方向
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[getNext(nums, n, fast)] > 0 {
			if slow == fast {
				if slow == getNext(nums, n, slow) { // 等于本身，退出继续寻找
					break
				}
				return true
			}
			slow = getNext(nums, n, slow)
			fast = getNext(nums, n, getNext(nums, n, fast))
		}
	}
	return false
}

func getNext(nums []int, n, cur int) int {
	return ((cur+nums[cur])%n + n) % n
}
