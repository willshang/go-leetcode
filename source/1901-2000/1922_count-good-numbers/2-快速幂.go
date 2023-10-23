package main

func main() {

}

var mod = int64(1000000007)

func countGoodNumbers(n int64) int {
	return int(mypow(5, (n+1)/2) * mypow(4, n/2) % mod)
}

func mypow(a int64, n int64) int64 {
	var res = int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n = n / 2
	}
	return res
}
