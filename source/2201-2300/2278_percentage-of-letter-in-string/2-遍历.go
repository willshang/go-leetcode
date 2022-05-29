package main

func main() {

}

// leetcode2278_字母在字符串中的百分比
func percentageLetter(s string, letter byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return 100 * count / len(s)
}
