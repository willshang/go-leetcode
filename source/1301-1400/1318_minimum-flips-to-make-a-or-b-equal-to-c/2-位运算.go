package main

func main() {

}

// leetcode1318_或运算的最小翻转次数
func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 31; i++ {
		A := a & 1
		B := b & 1
		C := c & 1
		if C == 0 {
			res = res + A + B
		} else {
			if A+B == 0 {
				res = res + 1
			}
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return res
}
