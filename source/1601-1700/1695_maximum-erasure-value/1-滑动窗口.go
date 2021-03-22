package main

func main() {

}

// leetcode1695_删除子数组的最大得分
func maximumUniqueSubarray(nums []int) int {
	res := 0
	sum := 0
	m := make(map[int]int)
	left := 0
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		sum = sum + nums[i]
		for m[nums[i]] > 1 {
			m[nums[left]]--
			sum = sum - nums[left]
			left++
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
