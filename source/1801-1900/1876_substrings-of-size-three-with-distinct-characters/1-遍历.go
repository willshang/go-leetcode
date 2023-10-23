package main

func main() {

}

// leetcode1876_长度为三且各字符不同的子字符串
func countGoodSubstrings(s string) int {
	res := 0
	for i := 2; i < len(s); i++ {
		if s[i-2] != s[i-1] && s[i-2] != s[i] && s[i-1] != s[i] {
			res++
		}
	}
	return res
}
