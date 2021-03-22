package main

func main() {

}

// leetcode1684_统计一致字符串的数目
func countConsistentStrings(allowed string, words []string) int {
	res := 0
	m := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = true
	}

	for _, word := range words {
		flag := true
		for i := 0; i < len(word); i++ {
			if _, ok := m[word[i]]; !ok {
				flag = false
				break
			}
		}
		if flag == true {
			res++
		}
	}
	return res
}
