package main

func main() {

}

func minOperations(nums []int, x int) int {
	n := len(nums)
	res := n + 1
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
	}
	if sum < x {
		return -1
	}
	if sum == x {
		return n
	}
	left := make([]int, n)
	right := make([]int, n)
	mLeft := make(map[int]int)
	mRight := make(map[int]int)
	left[0] = nums[0]
	mLeft[nums[0]] = 0
	right[n-1] = nums[n-1]
	mRight[nums[n-1]] = n - 1
	if left[0] == x || right[n-1] == x {
		return 1
	}
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + nums[i]
		mLeft[left[i]] = i
		if left[i] == x {
			res = min(res, i+1)
		}
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + nums[i]
		mRight[right[i]] = i
		if right[i] == x {
			res = min(res, n-i)
		}
	}
	for i := 1; i < n; i++ {
		left := left[i]
		if index, ok := mRight[x-left]; ok && i < index {
			target := n - (index - i - 1)
			res = min(res, target)
		}
	}
	if res == n+1 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
