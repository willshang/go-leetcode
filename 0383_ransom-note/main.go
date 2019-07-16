package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aaa", "aab"))
}
func canConstruct(ransomNote string, magazine string) bool {
	index := [26]int{}
	for i := 0; i < len(magazine); i++ {
		index[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		index[ransomNote[i]-'a']--
		if index[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
