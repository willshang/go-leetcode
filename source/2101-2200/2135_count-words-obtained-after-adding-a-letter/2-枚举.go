package main

func main() {

}

func wordCount(startWords []string, targetWords []string) int {
	m := make(map[int]bool)
	for i := 0; i < len(startWords); i++ {
		status := getStatus(startWords[i])
		m[status] = true
	}
	res := 0
	for i := 0; i < len(targetWords); i++ {
		status := getStatus(targetWords[i])
		for j := 0; j < len(targetWords[i]); j++ {
			v := int(targetWords[i][j] - 'a')
			if m[status^(1<<v)] == true { // 去除字符
				res++
				break
			}
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
