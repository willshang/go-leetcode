package main

func main() {

}

// leetcode1536_排布二进制网格的最少交换次数
func minSwaps(grid [][]int) int {
	n := len(grid)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				count++
			} else {
				break
			}
		}
		arr[i] = count
	}
	res := 0
	for i := 0; i < n-1; i++ {
		if arr[i] >= n-1-i { // 满足条件
			continue
		}
		j := i
		for ; j < n; j++ { // 往后找
			if arr[j] >= n-1-i {
				break
			}
		}
		if j == n { // 找不到
			return -1
		}
		for ; j > i; j-- { // 前移
			arr[j], arr[j-1] = arr[j-1], arr[j]
			res++
		}
	}
	return res
}
