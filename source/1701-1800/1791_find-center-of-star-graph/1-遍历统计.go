package main

func main() {

}

// leetcode1791_找出星型图的中心节点
func findCenter(edges [][]int) int {
	m := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		m[a]++
		m[b]++
	}
	for k, v := range m {
		if v == len(edges) {
			return k
		}
	}
	return -1
}
