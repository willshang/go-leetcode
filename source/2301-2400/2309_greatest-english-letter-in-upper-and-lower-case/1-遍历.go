package main

func main() {

}

// leetcode2309_兼具大小写的最好英文字母
func greatestLetter(s string) string {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}
	for i := 25; i >= 0; i-- {
		x := byte('a' + i)
		y := byte('A' + i)
		if m[x] == true && m[y] == true {
			return string(y)
		}
	}
	return ""
}
