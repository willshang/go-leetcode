package main

func main() {

}

// leetcode1812_判断国际象棋棋盘中一个格子的颜色
func squareIsWhite(coordinates string) bool {
	a := int(coordinates[0] - 'a')
	b := int(coordinates[1] - '1')
	return (a+b)%2 != 0
}
