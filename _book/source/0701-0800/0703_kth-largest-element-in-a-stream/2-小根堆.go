package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kth := Constructor(k, arr)
	kth.Add(3)
	kth.Add(10)
	fmt.Println(kth.k, kth.nums)
}

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	if k < len(nums) {
		sort.Ints(nums)
		nums = nums[len(nums)-k:]
	}
	// 向上调整
	Up(nums)
	return KthLargest{
		nums: nums,
		k:    k,
	}
}

func (k *KthLargest) Add(val int) int {
	if k.k > len(k.nums) {
		k.nums = append(k.nums, val)
		Up(k.nums)
	} else {
		if val > k.nums[0] {
			// 在堆顶，向下调整
			k.nums[0] = val
			Down(k.nums, 0)
		}
	}
	return k.nums[0]
}

func Down(nums []int, index int) {
	length := len(nums)
	minIndex := index
	for {
		left := 2*index + 1
		right := 2*index + 2
		if left < length && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < length && nums[right] < nums[minIndex] {
			minIndex = right
		}
		if minIndex == index {
			break
		}
		swap(nums, index, minIndex)
		index = minIndex
	}
}

func Up(nums []int) {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		minIndex := i
		left := 2*i + 1
		right := 2*i + 2
		if left < length && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < length && nums[right] < nums[minIndex] {
			minIndex = right
		}
		if i != minIndex {
			swap(nums, i, minIndex)
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
