package main

func main() {

}

// leetcode795_区间子数组个数
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	res := 0
	j := -1
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > right {
			j = i
		}
		if nums[i] >= left {
			count = i - j // 满足要求，如果大于right，则count=0
		}
		res = res + count
	}
	return res
}
