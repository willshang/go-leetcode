package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	//fmt.Println(buddyStrings("ad", "bc"))
	//fmt.Println(buddyStrings("a", "b"))
}

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return hasDouble(A)
	}
	first := -1
	second := -1
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return A[first] == B[second] && A[second] == B[first]
}

func hasDouble(s string) bool {
	seen := [26]int{}
	for i := range s {
		b := s[i] - 'a'
		if seen[b] >= 1 {
			return true
		}
		seen[b] = 1
	}
	return false
}
