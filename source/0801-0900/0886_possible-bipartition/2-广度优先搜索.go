package main

func main() {

}

// leetcode886_可能的二分法
func possibleBipartition(n int, dislikes [][]int) bool {
	m := make(map[int]int) // 分组： 0一组，1一组
	arr := make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		a, b := dislikes[i][0], dislikes[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	for i := 1; i <= n; i++ {
		// 没有被分配过，分配到0一组
		if _, ok := m[i]; ok == true {
			continue
		}
		m[i] = 0
		queue := make([]int, 0)
		queue = append(queue, i)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for i := 0; i < len(arr[node]); i++ {
				target := arr[node][i]
				if _, ok := m[target]; ok == false {
					m[target] = 1 - m[node] // 相反一组
					queue = append(queue, target)
				} else if m[node] == m[target] { // 已经分配，查看是否同一组
					return false
				}
			}
		}
	}
	return true
}
