package main

func main() {

}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		left := max(A[i][0], B[j][0])
		right := min(A[i][1], B[j][1])
		if left <= right {
			res = append(res, []int{left, right})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
