package main

func main() {

}

func isIdealPermutation(A []int) bool {
	if len(A) < 3 {
		return true
	}
	// 局部倒置首先是一个全局倒置，因此无需统计局部倒置
	// 只需要判断是否存在 0<=i<k<j<N 并且A[i]>A[j]
	maxValue := A[0] // 前2位数组最大值，如果存在maxValue > A[i]，则不会相等
	for i := 2; i < len(A); i++ {
		if maxValue > A[i] {
			return false
		}
		if maxValue < A[i-1] {
			maxValue = A[i-1]
		}
	}
	return true
}
