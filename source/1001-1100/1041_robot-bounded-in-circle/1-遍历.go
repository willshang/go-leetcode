package main

func main() {

}

// leetcode1041_困于环中的机器人
func isRobotBounded(instructions string) bool {
	var dx = []int{1, 0, -1, 0}
	var dy = []int{0, 1, 0, -1}
	dir := 0
	x, y := 0, 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'L' {
			dir = (dir + 3) % 4
		} else if instructions[i] == 'R' {
			dir = (dir + 1) % 4
		} else {
			x = x + dx[dir]
			y = y + dy[dir]
		}
	}
	return (x == 0 && y == 0) || dir != 0
}
