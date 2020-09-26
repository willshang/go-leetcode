package main

import "fmt"

func main() {
	// fmt.Println(isIsomorphic("foo", "baa")) // True
	// fmt.Println(isIsomorphic("ab", "aa")) // False
	// fmt.Println(isIsomorphic("aa", "ab")) // False
	fmt.Println(isIsomorphic("ab", "ca")) // True
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[int]int)
	n := make(map[int]int)

	for i := 0; i < len(s); i++ {
		a := int(s[i])
		b := int(t[i])
		if m[a] == 0 && n[b] == 0 {
			m[a] = b
			n[b] = a
		} else if m[a] != b || n[b] != a {
			return false
		}
	}
	return true
}
