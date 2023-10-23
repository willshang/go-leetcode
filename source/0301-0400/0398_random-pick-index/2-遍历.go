package main

import (
	"math/rand"
	"time"
)

func main() {

}

// leetcode398_随机数索引
type Solution struct {
	nums []int
	r    *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
		r:    rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (this *Solution) Pick(target int) int {
	res := 0
	count := 1
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			// 蓄水池采样法
			if this.r.Intn(count)+1 == count {
				res = i
			}
			count++
		}
	}
	return res
}
