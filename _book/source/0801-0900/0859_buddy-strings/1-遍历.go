package main

import "fmt"

func main() {
	//fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ad", "bc"))
	fmt.Println(buddyStrings("a", "b"))
}

// leetcode859_亲密字符串
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return hasDouble(A)
	}
	count := 2
	strA, strB := "", ""
	i := 0
	for ; count > 0 && i < len(A); i++ {
		if A[i] != B[i] {
			strA = string(A[i]) + strA
			strB = strB + string(B[i])
			count--
		}
	}
	return count == 0 && strA == strB && A[i:] == B[i:]
}

func hasDouble(s string) bool {
	seen := [26]bool{}
	for i := range s {
		b := s[i] - 'a'
		if seen[b] {
			return true
		}
		seen[b] = true
	}
	return false
}
