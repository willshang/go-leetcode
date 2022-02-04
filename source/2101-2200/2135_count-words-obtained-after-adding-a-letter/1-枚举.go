package main

func main() {

}

// leetcode2135_统计追加字母可以获得的单词数
func wordCount(startWords []string, targetWords []string) int {
	m := make(map[int]bool)
	for i := 0; i < len(startWords); i++ {
		status := getStatus(startWords[i])
		for j := 0; j < 26; j++ { // 枚举所有操作后的状态
			if (status>>j)&1 == 0 {
				m[status|(1<<j)] = true
			}
		}
	}
	res := 0
	for i := 0; i < len(targetWords); i++ {
		if m[getStatus(targetWords[i])] == true {
			res++
		}
	}
	return res
}

func getStatus(str string) int {
	status := 0
	for i := 0; i < len(str); i++ {
		v := int(str[i] - 'a')
		status = status | (1 << v)
	}
	return status
}
