package main

func main() {

}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		// arr := make([]int, m)

	}
	return res
}
