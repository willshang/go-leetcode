package main

func main() {

}

// leetcode2085_统计出现过一次的公共字符串
func countWords(words1 []string, words2 []string) int {
	a, b := make(map[string]int), make(map[string]int)
	for i := 0; i < len(words1); i++ {
		a[words1[i]]++
	}
	for i := 0; i < len(words2); i++ {
		b[words2[i]]++
	}
	res := 0
	for k, v := range a {
		if v == 1 && b[k] == 1 {
			res++
		}
	}
	return res
}
