package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// 程序员面试金典10.02_变位词组
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int)
	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		arr := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			arr[strs[i][j]-'a']++
		}
		if _, ok := m[arr]; ok {
			res[m[arr]] = append(res[m[arr]], strs[i])
		} else {
			m[arr] = len(res)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}
