package main

func main() {

}

func removeAnagrams(words []string) []string {
	flag := true
	for flag == true {
		for i := len(words) - 1; i >= 0; i-- {
			if i == 0 {
				flag = false
				break
			}
			if judge(words[i], words[i-1]) == true {
				words = append(words[:i], words[i+1:]...)
				flag = true
				break
			}
		}
	}
	return words
}

func judge(a, b string) bool {
	m := make(map[rune]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
