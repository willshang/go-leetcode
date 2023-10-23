package main

import "math"

func main() {

}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	n := len(routes)
	m := make(map[int][]int) // 该公交车经过第几条线路的数组
	dis := make([]int, n)    // source到第i条线路的距离
	arr := make([][]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]bool, n)
		dis[i] = -1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(routes[i]); j++ {
			node := routes[i][j]
			m[node] = append(m[node], i)
			for _, v := range m[node] {
				arr[i][v] = true
				arr[v][i] = true
			}
		}
	}
	queue := make([]int, 0)
	for _, v := range m[source] {
		dis[v] = 1
		queue = append(queue, v)
	}
	for len(queue) > 0 { // 广度优先计算source所在公交站台，到其它站台的最小距离
		node := queue[0]
		queue = queue[1:]
		for i := 0; i < n; i++ {
			if arr[node][i] == true && dis[i] == -1 { // node可以到i，且i未访问过
				dis[i] = dis[node] + 1
				queue = append(queue, i)
			}
		}
	}
	res := math.MaxInt32                  //  最短路径
	for i := 0; i < len(m[target]); i++ { // 遍历多终点，找最小
		v := m[target][i]
		if dis[v] != -1 && dis[v] < res {
			res = dis[v]
		}
	}
	if res < math.MaxInt32 {
		return res
	}
	return -1
}
