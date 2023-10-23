package main

func main() {

}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
	n, m := len(matrix), len(matrix[0])
	for a := 0; a < n; a++ { // 上边界
		arr := make([]int, m)    // 每列的和
		for b := a; b < n; b++ { // 下边界
			for j := 0; j < m; j++ {
				arr[j] = arr[j] + matrix[b][j]
			}
			temp := make(map[int]int)
			temp[0] = 1
			sum := 0
			for j := 0; j < m; j++ {
				sum = sum + arr[j]
				res = res + temp[sum-target]
				temp[sum]++
			}
		}
	}
	return res
}
