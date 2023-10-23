package main

func main() {

}

var mod int = 1337

func superPow(a int, b []int) int {
	a = a % mod
	if len(b) == 0 {
		return 1
	}
	res := 1
	for i := 0; i < len(b); i++ {
		res = mypow(res, 10) * mypow(a, b[i]) % mod
	}
	return res
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
