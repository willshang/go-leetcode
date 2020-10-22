package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	arr := make([]int, len(nums)+1)
	arr[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]%2
		arr[sum]++
		if sum >= k {
			res = res + arr[sum-k]
		}
	}
	return res
}
