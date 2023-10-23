package main

import "fmt"

func main() {
	fmt.Println(maxNiceDivisors(5))
}

// leetcode1808_好因子的最大数目
/*
由题意有: n=a1^k1 * a2^k2 *...*an^kn(如12=2^2 * 3^1)
其中
1、a1,a2,...,an是不同的质数(2,3不重复)
2、k1+k2+...+kn <=primeFactors
3、n的好因子,要能被每一个质因数(a1,a2,a3,...,an)整除，即好因子必须含有a1*a2*...*an作为因数
=>好因子的个数 k = k1*k2*...*kn =>求k最大，其中k1+..kn=primeFactors
等价于343题，整数拆分
*/

var mod = 1000000007

func maxNiceDivisors(primeFactors int) int {
	n := primeFactors
	if n <= 3 {
		return n
	}
	if n%3 == 0 {
		return pow(3, n/3) % mod
	} else if n%3 == 1 {
		return pow(3, (n-4)/3) * 4 % mod
	}
	return pow(3, n/3) * 2 % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b = b / 2
	}
	return res
}
