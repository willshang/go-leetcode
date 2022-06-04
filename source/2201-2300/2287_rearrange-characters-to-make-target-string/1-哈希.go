package main

func main() {

}

// leetcode2287_重排字符形成目标字符串
func rearrangeCharacters(s string, target string) int {
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m1[s[i]]++
	}
	for i := 0; i < len(target); i++ {
		m2[target[i]]++
	}
	res := len(s)
	for k, v := range m2 {
		value := m1[k] / v
		if value < res {
			res = value
		}
	}
	return res
}
