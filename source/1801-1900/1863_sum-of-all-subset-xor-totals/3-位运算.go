package main

func main() {

}

func subsetXORSum(nums []int) int {
	res := 0
	n := len(nums)
	total := 1 << n
	for i := 0; i < total; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				sum = sum ^ nums[j]
			}
		}
		res = res + sum
	}
	return res
}
