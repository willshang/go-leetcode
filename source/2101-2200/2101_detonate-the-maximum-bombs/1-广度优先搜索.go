package main

func main() {

}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && judge(bombs[i][0], bombs[i][1], bombs[j][0], bombs[j][1], bombs[i][2]) == true {
				arr[i] = append(arr[i], j)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		count := 1
		visited := make([]bool, n)
		queue := make([]int, 0)
		queue = append(queue, i)
		visited[i] = true
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for j := 0; j < len(arr[node]); j++ {
				next := arr[node][j]
				if visited[next] == false {
					count++
					queue = append(queue, next)
					visited[next] = true
				}
			}
		}
		if count > res {
			res = count
		}
	}
	return res
}

func judge(a, b, c, d, r int) bool {
	return r*r >= (a-c)*(a-c)+(b-d)*(b-d)
}
