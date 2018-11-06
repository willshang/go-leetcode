package main

import "fmt"

func main() {
	arr := [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}
	fmt.Println(flipAndInvertImage(arr))
}
func flipAndInvertImage(A [][]int) [][]int {
	for k := 0; k < len(A); k++{
		i,j := 0, len(A[k])-1
		for i < j{
			A[k][i],A[k][j] = invert(A[k][j]),invert(A[k][i])
			i++
			j--
		}
		if i == j {
			A[k][i] = invert(A[k][i])
		}
	}
	return A
}

func invert(i int)int  {
	if i == 0{
		return 1
	}
	return 0
}