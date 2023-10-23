package main

import "math/rand"

func main() {

}

// leetcode528_按权重随机选择
type Solution struct {
	nums  []int
	total int
}

func Constructor(w []int) Solution {
	total := 0
	arr := make([]int, len(w)) // 前缀和
	for i := 0; i < len(w); i++ {
		total = total + w[i]
		arr[i] = total
	}
	return Solution{
		nums:  arr,
		total: total,
	}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.total)
	left, right := 0, len(this.nums)
	for left < right {
		mid := left + (right-left)/2
		if this.nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
