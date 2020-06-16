package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	m := make(map[int32]int)
	arr := make([]int32, 0)
	res := 0
	for _, value := range s {
		if v, ok := m[value]; ok && v > 0 {
			for len(arr) > 0 && arr[0] != value {
				m[arr[0]]--
				arr = arr[1:]
			}
			m[arr[0]]--
			arr = arr[1:]
		}
		m[value]++
		arr = append(arr, value)
		if len(arr) > res {
			res = len(arr)
		}
	}
	return res
}
