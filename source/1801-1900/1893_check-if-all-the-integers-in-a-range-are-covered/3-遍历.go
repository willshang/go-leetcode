package main

func main() {

}

func isCovered(ranges [][]int, left int, right int) bool {
	for i := 0; i < len(ranges); i++ {
		a, b := ranges[i][0], ranges[i][1]
		if a <= left { // a到left已经覆盖：left移动到b后面
			left = b + 1
		}
		if b >= right { // right到b已经覆盖：right移动到a前面
			right = a - 1
		}
		if left > right {
			return true
		}
	}
	return false
}
