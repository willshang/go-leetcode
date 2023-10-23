package main

func main() {

}

// leetcode2260_必须拿起的最小连续卡牌数
func minimumCardPickup(cards []int) int {
	n := len(cards)
	res := n + 1
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		if v, ok := m[cards[i]]; ok && i-v+1 < res {
			res = i - v + 1
		}
		m[cards[i]] = i
	}
	if res == n+1 {
		return -1
	}
	return res
}
