package main

func main() {

}

// leetcode2027_转换字符串的最少操作次数
func minimumMoves(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			i = i + 2
			res++
		}
	}
	return res
}
