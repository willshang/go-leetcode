package main

func main() {

}

func isIdealPermutation(A []int) bool {
	if len(A) < 3 {
		return true
	}
	minValue := len(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < minValue {
			minValue = A[i]
		}
		if A[i-2] > minValue {
			return false
		}
	}
	return true
}
