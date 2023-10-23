package main

func main() {

}

// leetcode2108_找出数组中的第一个回文字符串
func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) == true {
			return words[i]
		}
	}
	return ""
}

func isPalindrome(a string) bool {
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}
