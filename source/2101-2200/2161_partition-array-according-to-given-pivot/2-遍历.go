package main

func main() {

}

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			res = append(res, nums[i])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] == pivot {
			res = append(res, nums[i])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > pivot {
			res = append(res, nums[i])
		}
	}
	return res
}
