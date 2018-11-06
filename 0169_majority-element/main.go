package main

import "fmt"

func main() {
	nums := []int{0,3,2,2,1,1,1,2,2,1,1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	ret, t := nums[0],1

	for i := 1; i < len(nums); i++{
		if ret == nums[i]{
			t++
		}else if t > 0{
			t--
		}else {
			ret = nums[i]
			t = 1
		}
	}
	return ret
}