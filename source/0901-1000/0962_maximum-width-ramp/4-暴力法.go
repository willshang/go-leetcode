package main

func main() {

}

func maxWidthRamp(A []int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i+res; j-- {
			if A[i] <= A[j] {
				res = max(res, j-i)
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
