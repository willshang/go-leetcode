package main

func main() {

}

// leetcode1941_检查是否所有字符出现次数相同
func areOccurrencesEqual(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	count := len(s) / len(m)
	for _, v := range m {
		if v != count {
			return false
		}
	}
	return true
}
