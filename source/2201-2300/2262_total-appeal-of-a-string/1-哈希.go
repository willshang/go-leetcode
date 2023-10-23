package main

import "fmt"

func main() {
	fmt.Println(appealSum("abbca"))
}

func appealSum(s string) int64 {
	res := int64(0)
	n := len(s)
	m := make(map[int]int)
	for i := 0; i < 26; i++ {
		m[i] = -1 // 上一次出现的位置
	}
	sum := 0
	for i := 0; i < n; i++ {
		v := int(s[i] - 'a')
		// 1、新字符没出现：在sum基础上加上(i+1)
		// 2、新字符出现过：在sum基础上加上i-m[v](到上一次出现的距离，只加后一段)
		sum = sum + (i - m[v])
		res = res + int64(sum)
		m[v] = i
	}
	return res
}
