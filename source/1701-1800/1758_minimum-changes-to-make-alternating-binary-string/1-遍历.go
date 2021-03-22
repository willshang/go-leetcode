package main

func main() {

}

// leetcode1758_生成交替二进制字符串的最少操作数
func minOperations(s string) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] == '1' { // a以偶0，奇1
			a++
		} else if i%2 == 1 && s[i] == '0' {
			a++
		}
		if i%2 == 0 && s[i] == '0' { // b以偶1，奇0
			b++
		} else if i%2 == 1 && s[i] == '1' {
			b++
		}
	}
	if a > b {
		return b
	}
	return a
}
