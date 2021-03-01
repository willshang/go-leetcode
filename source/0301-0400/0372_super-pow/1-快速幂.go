package main

func main() {

}

// leetcode372_超级次方
var mod int = 1337

func superPow(a int, b []int) int {
	a = a % mod
	if len(b) == 0 {
		return 1
	}
	x := mypow(a, b[len(b)-1])
	y := mypow(superPow(a, b[:len(b)-1]), 10)
	return x * y % mod
}

// a^n
func mypow(a int, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n = n / 2
	}
	return res
}
