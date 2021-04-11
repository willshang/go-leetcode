package main

func main() {

}

// leetcode403_青蛙过河
func canCross(stones []int) bool {
	n := len(stones)
	m := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		m[stones[i]] = make(map[int]int)
	}
	m[0][0] = 1
	for i := 0; i < n; i++ {
		for k := range m[stones[i]] {
			for next := k - 1; next <= k+1; next++ {
				if next > 0 {
					if _, ok := m[stones[i]+next]; ok {
						m[stones[i]+next][next] = 1
					}
				}
			}
		}
	}
	return len(m[stones[n-1]]) > 0
}
