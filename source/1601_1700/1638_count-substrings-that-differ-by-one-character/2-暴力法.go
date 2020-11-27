package main

func main() {

}

func countSubstrings(s string, t string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			length := j - i
			newStr := s[i:j]
			for k := 0; k <= len(t)-length; k++ {
				b := t[k : k+length]
				if compare(newStr, b) {
					res++
				}
			}
		}
	}
	return res
}

func compare(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
		if count >= 2 {
			return false
		}
	}
	if count == 0 {
		return false
	}
	return true
}
