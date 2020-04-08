package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividing(n int) bool {
	t := n
	for t > 0 {
		d := t % 10
		t = t / 10
		if d == 0 || n%d != 0 {
			return false
		}
	}
	return true
}
