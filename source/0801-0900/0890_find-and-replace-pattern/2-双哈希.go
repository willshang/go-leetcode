package main

func main() {

}

// leetcode890_查找和替换模式
func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		m1, m2 := make(map[byte]byte), make(map[byte]byte)
		flag := true
		for j := 0; j < len(pattern); j++ {
			x, y := pattern[j], words[i][j]
			if m1[x] == 0 && m2[y] == 0 {
				m1[x] = y
				m2[y] = x
			} else if (m1[x] == y && m2[y] == x) == false {
				flag = false
				break
			}
		}
		if flag == true {
			res = append(res, words[i])
		}
	}
	return res
}
