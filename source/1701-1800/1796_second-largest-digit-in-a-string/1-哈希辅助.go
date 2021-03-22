package main

func main() {

}

// leetcode1796_字符串中第二大的数字
func secondHighest(s string) int {
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			m[int(s[i]-'0')]++
		}
	}
	count := 0
	for i := 9; i >= 0; i-- {
		if m[i] > 0 {
			count++
			if count == 2 {
				return i
			}
		}
	}
	return -1
}
