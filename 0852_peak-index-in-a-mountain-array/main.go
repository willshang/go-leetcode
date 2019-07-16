package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
func peakIndexInMountainArray(A []int) int {
	l := len(A)
	for i := 1; i < l-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return 0
}

/*func peakIndexInMountainArray(A []int) int {
	l, r := 0, len(A)-1
	for{
		m := (l+r)/2
		switch  {
		case A[m] < A[m+1]:
			l = m
		case A[m-1] > A[m]:
			r = m
		default:
			return m
		}
	}
}*/
