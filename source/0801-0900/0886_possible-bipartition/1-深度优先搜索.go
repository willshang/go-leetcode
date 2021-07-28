package main

func main() {

}

var arr [][]int
var m map[int]int

func possibleBipartition(n int, dislikes [][]int) bool {
	m = make(map[int]int) // 分组： 0一组，1一组
	arr = make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		a, b := dislikes[i][0], dislikes[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	for i := 1; i <= n; i++ {
		// 没有被分配过，分配到0一组
		if _, ok := m[i]; ok == false && dfs(i, 0) == false {
			return false
		}
	}
	return true
}

func dfs(index int, value int) bool {
	if v, ok := m[index]; ok {
		return v == value // 已经分配，查看是否同一组
	}
	m[index] = value
	for i := 0; i < len(arr[index]); i++ {
		target := arr[index][i]
		if dfs(target, 1-value) == false { // 不喜欢的人，分配到对立组：1-value
			return false
		}
	}
	return true
}
