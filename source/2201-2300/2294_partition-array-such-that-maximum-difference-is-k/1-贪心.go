package main

import "sort"

func main() {

}

// leetcode2294_划分数组使最大差为K
func partitionArray(nums []int, k int) int {
	res := 1
	sort.Ints(nums)
	minValue := nums[0]
	for i := 0; i < len(nums); i++ {
		// 贪心：依次从小往大开始，把连续的、满足最大最小差值<=k划分为1组
		// 只需要考虑一组内的最大最小值，无需考虑顺序（划分为1组即可）
		if nums[i]-minValue > k {
			res++
			minValue = nums[i]
		}
	}
	return res
}
