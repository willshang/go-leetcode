package main

import "math/rand"

func main() {

}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	arr := make([]int, len(this.nums))
	copy(arr, this.nums)
	res := make([]int, len(this.nums))
	for i := 0; i < len(res); i++{
		j := rand.Intn(len(arr))
		res[i] = arr[j]
		arr = append(arr[:j], arr[j+1:]...)
	}
	return res
}
