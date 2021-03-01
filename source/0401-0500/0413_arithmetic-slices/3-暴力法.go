package main

func main() {

}

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0
	for i := 0; i < n-2; i++ {
		diff := A[i+1] - A[i]
		for j := i + 2; j < n; j++ {
			if A[j]-A[j-1] == diff {
				res++
			} else {
				break
			}
		}
	}
	return res
}
