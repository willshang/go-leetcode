package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	res := make([]int, len(A))
	for k := 0; k < len(A); k++ {
		if A[k]%2 == 0 {
			res[i] = A[k]
			i = i + 2
		} else {
			res[j] = A[k]
			j = j + 2
		}
	}
	return res
}
