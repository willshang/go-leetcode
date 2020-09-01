package main

import "fmt"

func main() {
	fmt.Println(getMaxLen([]int{-82, 62, 54, -63,
		-85, 53, -60, -59,
		29, 32, 59, -54,
		-29, -45}))
}

func getMaxLen(nums []int) int {
	arr := make([][]int, 0)
	temp := make([]int, 0)
	res := 0
	total := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			total = total * 1
		} else if nums[i] < 0 {
			total = total * -1
		} else {
			total = 0
		}
		if total > 0 {
			temp = append(temp, 1)
		} else if nums[i] == 0 {
			if len(temp) > 0 {
				arr = append(arr, temp)
				temp = make([]int, 0)
			}
			total = 1
		} else if total < 0 {
			temp = append(temp, -1)
		}
	}
	if len(temp) > 0 {
		arr = append(arr, temp)
	}
	for i := 0; i < len(arr); i++ {
		tempArr := arr[i]
		// 1、寻找最后一个1
		left, right := 0, len(tempArr)-1
		for 0 <= right && tempArr[right] != 1 {
			right--
		}
		res = max(res, right-left+1)
		// 2、寻找前后2个-1
		left, right = 0, len(tempArr)-1
		for left < len(tempArr) && tempArr[left] != -1 {
			left++
		}
		for 0 <= right && tempArr[right] != -1 {
			right--
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
