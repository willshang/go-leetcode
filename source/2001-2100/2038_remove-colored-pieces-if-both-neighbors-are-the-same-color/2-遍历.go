package main

func main() {

}

// leetcode2038_如果相邻两个颜色均相同则删除当前颜色
func winnerOfGame(colors string) bool {
	countA := 0
	countB := 0
	if len(colors) <= 2 {
		return false
	}
	for i := 2; i < len(colors); i++ {
		if colors[i-1] == colors[i-2] && colors[i-1] == colors[i] {
			if colors[i] == 'A' {
				countA++
			} else {
				countB++
			}
		}
	}
	if countA > countB {
		return true
	}
	return false
}
