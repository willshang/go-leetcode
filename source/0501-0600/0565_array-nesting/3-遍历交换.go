package main

func main() {

}

func arrayNesting(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		count := 1
		for nums[i] != i {
			count++
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
		if count > res {
			res = count
		}
	}
	return res
}
