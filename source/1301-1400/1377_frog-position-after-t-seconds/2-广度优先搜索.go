package main

func main() {

}

// leetcode1377_T秒后青蛙的位置
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	arr := make([][]int, n+1)
	res := make([]float64, n+1)
	res[1] = 1
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	visited := make([]bool, n+1) // 已经访问过的
	queue := make([]int, 0)
	queue = append(queue, 1)
	count := 0
	for len(queue) > 0 {
		length := len(queue)
		if count == t {
			break
		}
		for i := 0; i < length; i++ {
			start := queue[i]
			visited[start] = true
			count := 0
			for j := 0; j < len(arr[start]); j++ {
				next := arr[start][j]
				if visited[next] == false {
					count++
				}
			}
			if count == 0 {
				continue
			}
			per := res[start] / float64(count) // 每一跳的概率
			for j := 0; j < len(arr[start]); j++ {
				next := arr[start][j]
				if visited[next] == false {
					res[start] = res[start] - per // start-per
					res[next] = res[next] + per   // next+per
					queue = append(queue, next)
				}
			}
		}
		queue = queue[length:]
		count++
	}
	return res[target]
}
