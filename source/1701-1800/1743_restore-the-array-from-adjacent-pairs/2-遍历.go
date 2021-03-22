package main

func main() {

}

func restoreArray(adjacentPairs [][]int) []int {
	m := make(map[int][]int)
	for i := 0; i < len(adjacentPairs); i++ {
		a, b := adjacentPairs[i][0], adjacentPairs[i][1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	prev := 0
	cur := 0
	for k, v := range m {
		if len(v) == 1 {
			prev = k
			cur = v[0]
			break
		}
	}
	res := make([]int, 0)
	res = append(res, prev)
	n := len(adjacentPairs) + 1
	for n > len(res) {
		res = append(res, cur)
		for i := 0; i < len(m[cur]); i++ {
			if m[cur][i] != prev {
				prev = cur
				cur = m[cur][i]
				break
			}
		}
	}
	return res
}
