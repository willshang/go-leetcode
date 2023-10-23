package main

func main() {

}

func equalPairs(grid [][]int) int {
	res := 0
	n := len(grid)
	countMap := make(map[[200]int]int)
	for i := 0; i < n; i++ {
		arr := [200]int{}
		for j := 0; j < n; j++ {
			arr[j] = grid[i][j]
		}
		countMap[arr]++
	}
	for j := 0; j < n; j++ {
		arr := [200]int{}
		for i := 0; i < n; i++ {
			arr[i] = grid[i][j]
		}
		res = res + countMap[arr]
	}
	return res
}
