package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaoooo"))
	fmt.Println(checkInclusion("ab", "eidcaooab"))
}

// leetcode567_字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if arr1 == arr2 {
			return true
		}
		arr2[s2[i]-'a']--
		arr2[s2[i+len(s1)]-'a']++
	}
	return arr1 == arr2
}
