package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0{
		return 0
	}
	if length == 1{
		return nums[0]
	}

	temp := nums[0]
	max := nums[0]

	for i := 1; i < length; i++{
		if temp < 0{
			temp = nums[i]
		}else {
			temp = temp + nums[i]
		}


		if max < temp{
			max = temp
		}
	}
	return max
}