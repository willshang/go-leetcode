package main

func main() {

}

func kIncreasing(arr []int, k int) int {
	res := 0
	for i := 1; i <= k; i++ {
		num := make([]int, 0)
		num = append(num, arr[i-1])
		for j := i - 1; j < len(arr)-k; j = j + k {
			num = append(num, arr[j+k])
		}
		count := lengthOfLIS(num)
		res = res + len(num) - count
	}
	return res
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	arr := make([]int, 0)
	arr = append(arr, nums[0])
	for i := 1; i < len(nums); i++ {
		index := upperBound(arr, nums[i])
		if index == len(arr) {
			arr = append(arr, nums[i])
		} else {
			arr[index] = nums[i]
		}
	}
	return len(arr)
}

// 返回第一个大于等于target的位置
func upperBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 收缩左边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
