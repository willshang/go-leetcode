package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMinSwaps("6975", 8))
	//arr := []byte("241")
	//nextPermutation(arr)
	//fmt.Println(string(arr))
}

// leetcode1850_邻位交换的最小次数
func getMinSwaps(num string, k int) int {
	target := []byte(num)
	n := len(target)
	res := 0
	for i := 1; i <= k; i++ { // 求接下来的第k个排列
		nextPermutation(target)
	}
	// 统计交换次数
	arr := []byte(num)
	for i := 0; i < n; i++ {
		if arr[i] != target[i] {
			for j := i + 1; j < n; j++ {
				if arr[j] == target[i] { // 找到交换
					for k := j - 1; k >= i; k-- { // 把arr[j]交换到前面去
						arr[k], arr[k+1] = arr[k+1], arr[k]
						res++
					}
					break
				}
			}
		}
	}
	return res
}

// leetcode31.下一个排列
func nextPermutation(nums []byte) {
	n := len(nums)
	left := n - 2
	// 以12385764为例，从后往前找到5<7 的升序情况，目标值为左边的数5
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 { // 存在升序的情况
		right := n - 1
		// 从后往前，找到第一个大于目标值的数，如6>5，然后交换
		for right >= 0 && nums[right] <= nums[left] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	reverse(nums, left+1, n-1)
}

func reverse(nums []byte, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
