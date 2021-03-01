package main

func main() {

}

// leetcode775_全局倒置与局部倒置
func isIdealPermutation(A []int) bool {
	if len(A) < 3 {
		return true
	}
	for i := 0; i < len(A); i++ {
		if abs(i-A[i]) > 1 {
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
