package main

func main() {

}

// leetcode2044_统计按位或能得到最大值的子集数目
func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	total := 1 << n
	maxValue := 0
	res := 0
	for i := 0; i < total; i++ { // 枚举状态
		value := 0
		for j := 0; j < n; j++ { // 枚举该位
			if (i & (1 << j)) > 0 {
				value = value | nums[j]
			}
		}
		if value > maxValue {
			maxValue = value
			res = 1
		} else if value == maxValue {
			res++
		}
	}
	return res
}
