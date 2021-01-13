package main

func main() {

}

// leetcode1456_定长子串中元音的最大数目
func maxVowels(s string, k int) int {
	res := 0
	total := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) == true {
			total++
		}
		if i >= k {
			if isVowel(s[i-k]) == true {
				total--
			}
		}
		res = max(res, total)
	}
	return res
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' ||
		b == 'i' || b == 'o' || b == 'u'
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
