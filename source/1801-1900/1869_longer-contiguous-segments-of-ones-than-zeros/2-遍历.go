package main

func main() {

}

// leetcode1869_哪种连续子字符串更长
func checkZeroOnes(s string) bool {
	a, b := 0, 0
	aMax, bMax := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			a++
			b = 0
		} else if s[i] == '1' {
			a = 0
			b++
		}
		aMax = max(aMax, a)
		bMax = max(bMax, b)
	}
	return bMax > aMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
