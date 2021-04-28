package main

func main() {

}

// leetcode330_按要求补齐数组
func minPatches(nums []int, n int) int {
	res := 0
	target := 1
	for i := 0; target <= n; {
		// 对于正整数x，如果区间[1,x-1]内的所有数字都已经被覆盖，
		// 加上x后，则区间[1,2x-1]内的所有数字被覆盖。
		if i < len(nums) && nums[i] <= target {
			// 贪心思路：把当前在范围内的加上，target之前的[1,target-1]都已经满足
			target = target + nums[i]
			i++
		} else {
			// 没有就把target加入数组，范围扩大1倍
			target = target * 2
			res++
		}
	}
	return res
}
