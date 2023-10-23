package main

func main() {

}

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[i+1][j+1] = mat[i][j] + arr[i+1][j] + arr[i][j+1] - arr[i][j]
		}
	}
	res := min(n, m)
	left, right := 0, res // res可以为0
	for left <= right {
		mid := left + (right-left)/2
		if check(arr, mid, threshold) == true {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}
	return res
}

func check(arr [][]int, mid int, threshold int) bool {
	for i := 1; i+mid <= len(arr); i++ {
		for j := 1; j+mid <= len(arr[i]); j++ {
			sum := arr[i+mid-1][j+mid-1] - arr[i+mid-1][j-1] - arr[i-1][j+mid-1] + arr[i-1][j-1]
			if sum <= threshold {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
