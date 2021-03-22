package main

import "sort"

func main() {

}

// leetcode1657_确定两个字符串是否接近
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := make([]int, 26)
	arr2 := make([]int, 26)
	m1 := make(map[uint8]bool)
	m2 := make(map[uint8]bool)
	for i := 0; i < len(word1); i++ {
		arr1[word1[i]-'a']++
		arr2[word2[i]-'a']++
		m1[word1[i]-'a'] = true
		m2[word2[i]-'a'] = true
	}
	for key := range m1 {
		if m2[key] != true {
			return false
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
