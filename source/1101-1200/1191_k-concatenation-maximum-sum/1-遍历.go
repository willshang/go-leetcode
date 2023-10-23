package main

func main() {

}

// leetcode1191_K次串联后最大子数组之和
func kConcatenationMaxSum(arr []int, k int) int {
	cur := 0
	sum := 0
	res := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		cur = max(cur+arr[i], arr[i])
		res = max(res, cur)
	}
	for i := 0; i < len(arr); i++ {
		cur = max(cur+arr[i], arr[i])
		res = max(res, cur)
	}
	if sum <= 0 {
		return res % 1000000007
	}
	return (res + (k-2)*sum) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
