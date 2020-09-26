package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		arr := []byte(strs[i])
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		newStr := string(arr)
		if _, ok := m[newStr]; ok {
			res[m[newStr]] = append(res[m[newStr]], strs[i])
		} else {
			m[newStr] = len(res)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}
