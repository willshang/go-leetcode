package main

import (
	"fmt"
)

func main() {
	fmt.Println(15, 14, 13, 12, 11, 10)
	fmt.Println(subStrHash("xmmhdakfursinye", 2, 45, 3, 21))
}

// leetcode2156_查找给定哈希值的子串
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)
	p := 1   // power的k次方
	sum := 0 // 和
	// 第一个长度为k的结果
	for i := n - 1; i >= n-k; i-- {
		p = p * power % modulo
		a := int(s[i]-'a') + 1
		sum = (a + sum*power) % modulo
	}
	res := s[n-k:] // 题目保证至少有1个
	// 逆序计算：滑动窗口
	for i := n - k - 1; i >= 0; i-- {
		a := int(s[i]-'a') + 1 + sum*power      // k+1位
		b := (int(s[i+k]-'a') + 1) * p % modulo // s[i+k]*power^k
		// 模p运算：分配律
		// (a+b) % mod  = (a%mod + b%mod) % mod
		// (a-b) % mod  = (a%mod - b%mod ) % mod
		// (a-b) % mod  = (a%mod - b%mod + mod ) % mod：go有负数，+mod再mod
		sum = (a - b + modulo) % modulo
		if sum == hashValue {
			res = s[i : i+k]
		}
	}
	return res
}
