package main

func main() {

}

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	res := make([]int, 0)
	temp := nums[0]
	for i := 1; i < n; i++ {
		temp = temp ^ nums[i]
	}
	target := 1<<maximumBit - 1
	for i := n - 1; i >= 0; i-- {
		res = append(res, temp^target)
		temp = temp ^ nums[i]
	}
	return res
}
