package main

func main() {

}

func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	if 2*k+1 <= n {
		sum := 0
		for i := 0; i < 2*k+1; i++ {
			sum = sum + nums[i]
		}
		for i := k; i < n-k; i++ {
			if i != k {
				sum = sum + nums[i+k] - nums[i-k-1]
			}
			res[i] = sum / (2*k + 1)
		}
	}
	return res
}
