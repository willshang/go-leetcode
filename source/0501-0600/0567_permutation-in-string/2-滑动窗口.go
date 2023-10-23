package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaoooo"))
	fmt.Println(checkInclusion("ab", "eidcaooab"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if compare(m1, m2) {
			return true
		}
		m2[s2[i]-'a']--
		if m2[s2[i]-'a'] == 0 {
			delete(m2, s2[i]-'a')
		}
		m2[s2[i+len(s1)]-'a']++
	}
	return compare(m1, m2)
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
