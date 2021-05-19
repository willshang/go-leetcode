package main

func main() {

}

func subsetXORSum(nums []int) int {
	res := 0
	n := len(nums)
	left := 1 << n
	right := 1 << (n + 1)
	for i := left; i < right; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				sum = sum ^ nums[j]
			}
		}
		res = res + sum
	}
	return res
}
