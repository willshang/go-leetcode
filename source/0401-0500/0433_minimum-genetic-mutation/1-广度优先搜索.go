package main

import "fmt"

func main() {
	fmt.Println(minMutation("AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC",
		"AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}))
}

// leetcode433_最小基因变化
func minMutation(start string, end string, bank []string) int {
	arr := []byte{'A', 'T', 'C', 'G'}
	m := make(map[string]bool)
	for i := 0; i < len(bank); i++ {
		m[bank[i]] = true
	}
	if _, ok := m[end]; ok == false {
		return -1
	}
	res := 0
	queue := make([]string, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			str := queue[i]
			for j := 0; j < len(str); j++ {
				for k := 0; k < len(arr); k++ {
					if arr[k] != str[j] {
						newStr := str[:j] + string(arr[k]) + str[j+1:]
						if _, ok := m[newStr]; ok {
							queue = append(queue, newStr)
							delete(m, newStr)
						}
						if newStr == end {
							return res
						}
					}
				}
			}
		}
		queue = queue[length:]
	}
	return -1
}
