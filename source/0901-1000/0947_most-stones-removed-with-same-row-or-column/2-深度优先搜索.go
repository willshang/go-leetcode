package main

func main() {

}

var arr [][]int
var m []bool

func removeStones(stones [][]int) int {
	n := len(stones)
	arr = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stones[i][0] == stones[j][0] || // 同行
				stones[i][1] == stones[j][1] { // 同列
				arr[i] = append(arr[i], j)
				arr[j] = append(arr[j], i)
			}
		}
	}
	m = make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		count = count + dfs(i)
	}
	return len(stones) - count
}

func dfs(index int) int {
	if m[index] == true {
		return 0
	}
	m[index] = true
	for i := 0; i < len(arr[index]); i++ {
		dfs(arr[index][i])
	}
	return 1
}
