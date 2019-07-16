package main

import "fmt"

func main() {
	fmt.Println(countPrimes(100000))
}

/*
//超时
func countPrimes(n int) int {
	isDeleted := make(map[int]bool)
	count := 0
	for i := 2; i < n; i++{
		if !isDeleted[i]{
			count++
			for j := 1; i * j <= n; j++{
				isDeleted[i*j] = true
			}
		}
	}
	return count
}*/

//大于2的偶数都不是质数，所以我们只需要判断奇数是不是质数。
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	isComposite := make([]bool, n)
	count := n / 2
	//奇数
	for i := 3; i*i < n; i += 2 {
		fmt.Println(i, count)
		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j = j + 2*i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}

	}
	return count
}
