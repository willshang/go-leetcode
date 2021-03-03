package main

func main() {

}

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		m1, m2 := make(map[byte]byte), make(map[byte]byte)
		flag := true
		for j := 0; j < len(pattern); j++ {
			x, y := pattern[j], words[i][j]
			a, ok1 := m1[x]
			b, ok2 := m2[y]
			if ok1 == false && ok2 == false {
				m1[x] = y
				m2[y] = x
			}
			if (ok1 == true && ok2 == false) || (ok1 == false && ok2 == true) ||
				(ok1 == true && ok2 == true && (a != y || b != x)) {
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
