package main

func main() {

}

// leetcode1763_最长的美好子字符串
func longestNiceSubstring(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if judge(s[i:j]) == true && len(s[i:j]) > len(res) {
				res = s[i:j]
			}
		}
	}
	return res
}

func judge(str string) bool {
	a := [26]int{}
	A := [26]int{}
	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' {
			a[str[i]-'a']++
		} else {
			A[str[i]-'A']++
		}
	}
	for i := 0; i < 26; i++ {
		if (a[i] > 0 && A[i] == 0) || (a[i] == 0 && A[i] > 0) {
			return false
		}
	}
	return true
}
