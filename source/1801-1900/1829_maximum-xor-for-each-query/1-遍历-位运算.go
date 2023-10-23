package main

func main() {

}

// leetcode1829_每个查询的最大异或值
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	res := make([]int, n)
	temp := nums[0]
	res[n-1] = temp
	for i := 1; i < n; i++ {
		temp = temp ^ nums[i]
		res[n-1-i] = temp
	}
	target := 1<<maximumBit - 1
	for i := 0; i < n; i++ {
		res[i] = res[i] ^ target
	}
	return res
}
