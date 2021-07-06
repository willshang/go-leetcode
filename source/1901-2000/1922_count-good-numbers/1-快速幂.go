package main

func main() {

}

// leetcode1922_统计好数字的数目
var mod = int64(1000000007)

func countGoodNumbers(n int64) int {
	res := mypow(20, n/2)
	if n%2 == 1 {
		res = res * 5 % mod
	}
	return int(res)
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
