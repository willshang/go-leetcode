package main

import "fmt"

func main() {
	s := "abcdddacbeeebcd"
	p := "bcd"
	fmt.Println(findAnagrams(s, p))
}

// leetcode438_找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		arr1[p[i]-'a']++
		arr2[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if arr1 == arr2 {
			res = append(res, i)
		}
		arr2[s[i]-'a']--
		arr2[s[i+len(p)]-'a']++
	}
	if arr1 == arr2 {
		res = append(res, len(s)-len(p))
	}
	return res
}
