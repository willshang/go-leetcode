package main

func main() {

}

// leetcode986_区间列表的交集
func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i][0] <= B[j][1] && B[j][0] <= A[i][1] {
			left := max(A[i][0], B[j][0])
			right := min(A[i][1], B[j][1])
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
