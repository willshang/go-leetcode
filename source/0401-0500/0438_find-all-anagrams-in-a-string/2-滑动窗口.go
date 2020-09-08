package main

import "fmt"

func main() {
	s := "abcdddacbeeebcd"
	p := "bcd"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m1[p[i]-'a']++
		m2[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if compare(m1, m2) {
			res = append(res, i)
		}
		m2[s[i]-'a']--
		if m2[s[i]-'a'] == 0 {
			delete(m2, s[i]-'a')
		}
		m2[s[i+len(p)]-'a']++
	}
	if compare(m1, m2) {
		res = append(res, len(s)-len(p))
	}
	return res
}

func compare(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}
