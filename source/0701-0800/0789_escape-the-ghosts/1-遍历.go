package main

func main() {

}

// leetcode789_逃脱阻碍者
func escapeGhosts(ghosts [][]int, target []int) bool {
	a := abs(target[0]) + abs(target[1])
	for i := 0; i < len(ghosts); i++ {
		b := abs(ghosts[i][0]-target[0]) + abs(ghosts[i][1]-target[1])
		if b <= a {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
