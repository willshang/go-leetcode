package main

func main() {

}

// leetcode554_砖墙
func leastBricks(wall [][]int) int {
	maxCount := 0
	m := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		index := 0
		for j := 0; j < len(wall[i])-1; j++ {
			index = index + wall[i][j]
			m[index]++ // 保留去除开头和结尾的位置(空隙地方)
			if maxCount <= m[index] {
				maxCount = m[index]
			}
		}
	}
	return len(wall) - maxCount
}
