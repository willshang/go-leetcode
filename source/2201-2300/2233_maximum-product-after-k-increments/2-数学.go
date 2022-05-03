package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{0, 4}, 5))
}

var mod = 1000000007

func maximumProduct(nums []int, k int) int {
	res := 1
	sort.Ints(nums)
	n := len(nums)
	nums = append(nums, math.MaxInt32/100)
	for i := 0; i < n; i++ {
		sum := (i + 1) * (nums[i+1] - nums[i])
		if k >= sum {
			k = k - sum
			continue
		}
		b := k / (i + 1) // 加多少
		d := k % (i + 1) // 剩下多少+1
		for j := 0; j <= i; j++ {
			nums[j] = nums[i] + b // 在nums[i]基础上加
			if j < d {
				nums[j]++
			}
		}
		break
	}
	for i := 0; i < n; i++ {
		res = res * nums[i] % mod
	}
	return res
}
