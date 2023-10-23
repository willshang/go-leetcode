package main

import "fmt"

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
}

// leetcode985_查询后的偶数和
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, 0)
	sum := 0
	for _, value := range A {
		if value%2 == 0 {
			sum = sum + value
		}
	}
	for i := 0; i < len(queries); i++ {
		value := queries[i][0]
		index := queries[i][1]
		if A[index]%2 == 0 {
			sum = sum - A[index]
		}
		A[index] = A[index] + value
		if A[index]%2 == 0 {
			sum = sum + A[index]
		}
		res = append(res, sum)
	}
	return res
}
