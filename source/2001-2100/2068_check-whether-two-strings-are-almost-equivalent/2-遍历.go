package main

func main() {

}

// leetcode2068_检查两个字符串是否几乎相等
func checkAlmostEquivalent(word1 string, word2 string) bool {
	a := [26]int{}
	for i := 0; i < len(word1); i++ {
		a[int(word1[i]-'a')]++
	}
	for i := 0; i < len(word2); i++ {
		a[int(word2[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if a[i] > 3 || a[i] < -3 {
			return false
		}
	}
	return true
}
