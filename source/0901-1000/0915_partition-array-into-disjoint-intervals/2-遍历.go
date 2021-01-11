package main

func main() {

}

// leetcode915_分割数组
func partitionDisjoint(A []int) int {
	n := len(A)
	res := 0
	maxLeft := A[0]
	maxValue := A[0]
	for i := 1; i < n; i++ {
		maxValue = max(maxValue, A[i])
		if A[i] < maxLeft {
			res = i
			maxLeft = maxValue
		}
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
