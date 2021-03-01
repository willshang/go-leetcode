package main

func main() {

}

func maxSubarraySumCircular(A []int) int {
	n := len(A)
	res := A[0]
	sum := 0
	for i := 0; i < n; i++ {
		if sum > 0 {
			sum = sum + A[i]
		} else {
			sum = A[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
