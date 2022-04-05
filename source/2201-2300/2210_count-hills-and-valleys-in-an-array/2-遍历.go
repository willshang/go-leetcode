package main

func main() {

}

// leetcode2210_统计数组中峰和谷的数量
func countHillValley(nums []int) int {
	res := 0
	n := len(nums)
	flag := 0 // 1：递增，2：递减
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			if flag == 2 {
				res++
			}
			flag = 1
		} else if nums[i-1] > nums[i] {
			if flag == 1 {
				res++
			}
			flag = 2
		}
	}
	return res
}
