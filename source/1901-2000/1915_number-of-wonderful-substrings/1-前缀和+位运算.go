package main

import "fmt"

func main() {
	fmt.Println(wonderfulSubstrings("abcdef"))
}

// leetcode1915_最美子字符串的数目
func wonderfulSubstrings(word string) int64 {
	res := 0
	m := make(map[int]int)
	m[0] = 1 // 初始状态没有任何字符，相当于全是偶数
	cur := 0
	for i := 0; i < len(word); i++ {
		value := int(word[i] - 'a')
		cur = cur ^ (1 << value) // 位运算-异或后的状态结果，表示到当前下标各字母出现的奇偶次数
		res = res + m[cur]       // 跟当前状态相同
		for i := 0; i < 10; i++ {
			count := m[cur^(1<<i)] // 跟当前状态差1位相同
			res = res + count
		}
		m[cur]++
	}
	return int64(res)
}
