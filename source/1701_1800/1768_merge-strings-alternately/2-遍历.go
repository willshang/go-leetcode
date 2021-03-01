package main

func main() {

}

// leetcode1768_交替合并字符串
func mergeAlternately(word1 string, word2 string) string {
	res := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			res = res + string(word1[i])
		}
		if i < len(word2) {
			res = res + string(word2[i])
		}
	}
	return res
}
