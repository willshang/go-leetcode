package main

func main() {

}

func minSteps(s string, t string) int {
	res := 0
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; ok {
			m[t[i]]--
		} else {
			m[t[i]] = -1
		}
	}
	for _, v := range m {
		if v > 0 {
			res = res + v
		}
	}
	return res
}
