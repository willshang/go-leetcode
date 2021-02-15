package main

func main() {

}

// leetcode1754_构造字典序最大的合并字符串
func largestMerge(word1 string, word2 string) string {
	res := make([]byte, len(word1)+len(word2))
	for i := 0; i < len(res); i++ {
		if word1 < word2 {
			res[i], word2 = word2[0], word2[1:]
		} else {
			res[i], word1 = word1[0], word1[1:]
		}
	}
	return string(res)
}
