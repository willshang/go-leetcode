package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 1; i <= rowIndex; i++ {
		res = append(res, 1)
		for j := i - 1; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}

	}
	return res
}
