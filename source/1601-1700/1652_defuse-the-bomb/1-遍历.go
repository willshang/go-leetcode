package main

func main() {

}

// leetcode1652_拆炸弹
func decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, n)
	if k == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		sum := 0
		for j := 1; j <= abs(k); j++ {
			if k > 0 {
				sum = sum + code[(i+j+n)%n]
			} else {
				sum = sum + code[(i-j+n)%n]
			}
		}
		res[i] = sum
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
