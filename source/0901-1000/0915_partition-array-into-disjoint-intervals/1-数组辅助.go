package main

func main() {

}

func partitionDisjoint(A []int) int {
	n := len(A)
	maxLeft := make([]int, n)
	minRight := make([]int, n)
	maxValue := A[0]
	for i := 0; i < n; i++ {
		if maxValue < A[i] {
			maxValue = A[i]
		}
		maxLeft[i] = maxValue
	}
	minValue := A[n-1]
	for i := n - 1; i >= 0; i-- {
		if minValue > A[i] {
			minValue = A[i]
		}
		minRight[i] = minValue
	}
	for i := 1; i < n; i++ {
		if maxLeft[i-1] <= minRight[i] {
			return i
		}
	}
	return -1
}
