package main

func main() {

}

func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)
	for i := 0; i < n; i++ {
		res[i] = nums[i]
		res[i+n] = nums[i]
	}
	return res
}
