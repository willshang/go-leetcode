package main

import "fmt"

func main() {

}

func countPalindromicSubsequence(s string) int {
	n := len(s)
	arr := [26][]int{}
	for i := 0; i < n; i++ {
		index := int(s[i] - 'a')
		arr[index] = append(arr[index], i)
	}
	m := make(map[string]bool)
	for i := 0; i < 26; i++ { // 枚举2边字符
		if len(arr[i]) <= 1 {
			continue
		}
		left, right := arr[i][0], arr[i][len(arr[i])-1]
		for j := 0; j < 26; j++ {
			if len(arr[j]) == 0 {
				continue
			}
			if i == j && len(arr[i]) > 2 {
				m[fmt.Sprintf("%d,%d,%d", i, i, i)] = true
				continue
			}
			flag := false
			for k := 0; k < len(arr[j]); k++ {
				if left < arr[j][k] && arr[j][k] < right {
					flag = true
					break
				}
			}
			if flag == true {
				m[fmt.Sprintf("%d,%d,%d", i, j, i)] = true
			}
		}
	}
	return len(m)
}
