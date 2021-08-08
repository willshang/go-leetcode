package main

func main() {

}

// leetcode1955_统计特殊子序列的数目
var mod = 1000000007

func countSpecialSubsequences(nums []int) int {
	n := len(nums)
	a, b, c := 0, 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			a = (a*2 + 1) % mod // 之前0组合+之前0组合加上当前0+单独0
		} else if nums[i] == 1 {
			b = (b*2 + a) % mod // 之前01组合+之前01组合加上当前1+之前0组合加上当前1
		} else if nums[i] == 2 {
			c = (c*2 + b) % mod // 之前012组合+之前012组合加上当前值2+之前01组合加上当前2
		}
	}
	return c
}
