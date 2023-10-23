package main

func main() {

}

// leetcode2039_网络空闲的时刻
func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	arr := make([][]int, n) // 邻接表：i=>j的集合
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1] // a=>b
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	res := 0
	visited := make([]bool, n)
	visited[0] = true
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		a, dis := node[0], node[1]
		if a != 0 {
			total := (2*dis-1)/patience[a]*patience[a] + 2*dis
			res = max(res, total)
		}
		for i := 0; i < len(arr[a]); i++ {
			b := arr[a][i]
			if visited[b] == false {
				queue = append(queue, [2]int{b, dis + 1})
				visited[b] = true
			}
		}
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
