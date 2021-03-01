package main

func main() {

}

// leetcode1023_驼峰式匹配
func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		str := queries[i]
		count := 0
		flag := true
		for j := 0; j < len(str); j++ {
			if count < len(pattern) && str[j] == pattern[count] {
				count++
				continue
			}
			if 'A' <= str[j] && str[j] <= 'Z' {
				flag = false
				break
			}
		}
		if flag == true && count == len(pattern) {
			res[i] = true
		}
	}
	return res
}
