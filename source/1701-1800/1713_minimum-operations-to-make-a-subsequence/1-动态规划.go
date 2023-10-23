package main

func main() {

}

// leetcode1713_得到子序列的最少操作次数
func minOperations(target []int, arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(target); i++ {
		m[target[i]] = i
	}
	nums := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if v, ok := m[arr[i]]; ok {
			nums = append(nums, v)
		}
	}
	res := lengthOfLIS(nums) // leetcode 300.最长上升子序列
	return len(target) - res
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	arr := make([]int, len(nums)+1)
	arr[1] = nums[0]
	res := 1
	for i := 1; i < len(nums); i++ {
		if arr[res] < nums[i] {
			res++
			arr[res] = nums[i]
		} else {
			left, right := 1, res
			index := 0
			for left <= right {
				mid := left + (right-left)/2
				if arr[mid] < nums[i] {
					index = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			arr[index+1] = nums[i]
		}
	}
	return res
}
