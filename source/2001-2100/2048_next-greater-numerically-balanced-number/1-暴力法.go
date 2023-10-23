package main

import "strconv"

func main() {

}

// leetcode2048_下一个更大的数值平衡数
func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if judge(i) == true {
			return i
		}
	}
}

func judge(a int) bool {
	s := strconv.Itoa(a)
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'0')]++
	}
	for k, v := range m {
		if k != v {
			return false
		}
	}
	return true
}
