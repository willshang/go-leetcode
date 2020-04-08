package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
}

func binaryGap(N int) int {
	A := make([]int, 0)
	index := 0
	for N > 0 {
		if N%2 == 1 {
			A = append(A, index)
		}
		index++
		N = N / 2
	}
	res := 0
	for i := 0; i < len(A)-1; i++ {
		res = max(res, A[i+1]-A[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*func binaryGap(N int) int {
	res := 0
	gap := 0

	for N > 0{
		fmt.Println(N)
		if N&1 == 1{
			res = max(res,gap)
			gap = 1
		}else if gap > 0{
			gap++
		}
		N = N / 2
	}
	return res
}

func max(a,b int) int  {
	if a > b{
		return a
	}
	return b
}*/
