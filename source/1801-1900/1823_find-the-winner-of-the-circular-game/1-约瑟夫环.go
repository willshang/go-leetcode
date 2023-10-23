package main

func main() {

}

// leetcode1823_找出游戏的获胜者
func findTheWinner(n int, k int) int {
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + k) % i
	}
	return idx + 1
}
